package db

import (
	"context"
	"fmt"
	"komo/lib/engine"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

type Postgres struct {
	Conn   *pgx.Conn
	DbName string
	Uri    string
	Logger *zap.Logger
}

var Logger *zap.Logger

func (p *Postgres) QueryRowBg(sql string, arguments ...any) *engine.Result[pgx.Row] {
	res := engine.NewResult[pgx.Row]()
	row := p.Conn.QueryRow(context.Background(), sql, arguments...)

	res.WithData(&row)
	return res
}

func (p *Postgres) QueryBg(sql string, arguments ...any) *engine.Result[pgx.Rows] {
	res := engine.NewResult[pgx.Rows]()
	rows, err := p.Conn.Query(context.Background(), sql, arguments...)
	if err != nil {
		res.WithError(err)
	}

	res.WithData(&rows)
	return res
}

func (p *Postgres) ExecBg(sql string, arguments ...any) *engine.Result[pgconn.CommandTag] {
	res := engine.NewResult[pgconn.CommandTag]()
	cmd, err := p.Conn.Exec(context.Background(), sql, arguments...)
	if err != nil {
		res.WithError(err)
	}

	res.WithData(&cmd)
	return res
}

func (p *Postgres) Connect() {
	conn, err := pgx.Connect(context.Background(), p.Uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Db.Conn = conn
}

func (p *Postgres) Close() {
	p.Conn.Close(context.Background())
}

var Db *Postgres

func extractNameFromUri(uri string) string {
	dbUrl, err := pgx.ParseConfig(uri)
	if err != nil {
		return ""
	}
	return dbUrl.Database
}

func NewPg(uri string) {
	dbName := extractNameFromUri(uri)

	Db = &Postgres{
		Uri:    uri,
		DbName: dbName,
		Logger: engine.Logger.With(
			zap.String("postgres", dbName),
		),
	}

	Logger = Db.Logger

	Db.Connect()
	if err := Db.Ping(); !err.IsOk() {
		engine.Logger.Fatal(err.Error.Error())
	}
}

func (p *Postgres) Ping() *engine.Result[bool] {
	result := engine.NewResult(true)

	if err := p.Conn.Ping(context.Background()); err != nil {
		p.Logger.Error(err.Error())

		return result.WithError(err)
	}

	p.Logger.Info("DB connected")

	return result.WithPureData(true)
}
