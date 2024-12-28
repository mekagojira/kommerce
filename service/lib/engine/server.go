package engine

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"gopkg.in/validator.v2"
)

var Server = chi.NewRouter()

type Ctx[I any, O any] struct {
	Url    string
	Id     string
	Req    *I
	Logger *zap.Logger
	w      http.ResponseWriter
	r      *http.Request
}

type Response[T any] struct {
	Code    string `json:"code"`
	Data    *T     `json:"data"`
	Message string `json:"message"`
	status  int
}

const (
	CodeOk    = "OK"
	CodeError = "ERROR"
)

func (ctx *Ctx[I, O]) Ok(data *O) *Response[O] {
	return &Response[O]{
		Code:    CodeOk,
		Data:    data,
		Message: "",
		status:  200,
	}
}

func (ctx *Ctx[I, O]) Error(code string, message string) *Response[O] {
	return &Response[O]{
		Code:    code,
		Data:    nil,
		Message: message,
		status:  400,
	}
}

func (ctx *Ctx[I, O]) ServerError(message ...string) *Response[O] {
	res := ctx.Error(CodeError, "Something went wrong")
	if len(message) > 0 {
		return ctx.Error(CodeError, message[0])
	}
	res.status = 500
	return res
}

func (ctx *Ctx[I, O]) BadRequest(message ...string) *Response[O] {
	if len(message) > 0 {
		return ctx.Error(CodeError, message[0])
	}
	return ctx.Error(CodeError, "Bad request")
}

func (ctx *Ctx[I, O]) SendResponse(res *Response[O]) {
	ctx.w.Header().Add("Content-Type", "application/json")
	ctx.w.WriteHeader(res.status)
	json.NewEncoder(ctx.w).Encode(res)
}

func (ctx *Ctx[I, O]) ParseRequest() *Result[I] {
	var req I
	var result Result[I]

	if err := json.NewDecoder(ctx.r.Body).Decode(&req); err != nil {
		return result.WithError(err)
	}

	if err := validator.Validate(req); err != nil {
		return result.WithError(err)
	}

	return result.WithData(&req)
}

func RegisterEndpoint[I any, O any](url string, handler func(*Ctx[I, O]) *Response[O]) {
	Logger.Info("Registering endpoint", zap.String("URL", url))

	Server.Post(url, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Ctx[I, O]{
			Url: url,
			r:   r,
			w:   w,
		}
		if err := GetUid(); err.Error != nil {
			ctx.SendResponse(ctx.ServerError())
			return
		} else {
			ctx.Id = *err.Data
			ctx.Logger = Logger.With(zap.String("id", ctx.Id))
		}

		ctx.Logger.Info(url)

		if result := ctx.ParseRequest(); !result.IsOk() {
			ctx.SendResponse(ctx.BadRequest())
			return
		} else {
			ctx.Req = result.Data
		}

		ctx.SendResponse(handler(ctx))
	})
}

func StartServer(port string) {
	log.Println("Starting server on port " + port)

	http.ListenAndServe(":"+port, Server)
}
