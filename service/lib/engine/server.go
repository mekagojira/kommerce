package engine

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"gopkg.in/validator.v2"
)

var Server = chi.NewRouter()

func init() {
	Server.Use(middleware.Timeout(15 * time.Second))
}

type CtxUser struct {
	UserId string
	Role   string
}

type Ctx[I any, O any] struct {
	Url    string
	Id     string
	Req    *I
	Logger *zap.Logger
	w      http.ResponseWriter
	r      *http.Request
	User   CtxUser
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

func (ctx *Ctx[I, O]) Unauthorized(message ...string) *Response[O] {
	res := ctx.Error(CodeError, "Unauthorized")

	if len(message) > 0 {
		res = ctx.Error(CodeError, message[0])
	}
	res.status = 401
	return res
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

func RegisterEndpoint[I any, O any](url string, handler func(*Ctx[I, O]) *Response[O], roles ...string) {
	Logger.Info("Registering endpoint", zap.String("URL", url))

	Server.Post(url, func(w http.ResponseWriter, r *http.Request) {
		author := r.Header.Get("Authorization")

		ctx := &Ctx[I, O]{
			Url: url,
			r:   r,
			w:   w,
		}

		if res := GetUid(); IsError(res) {
			ctx.SendResponse(ctx.ServerError())
			return
		} else {
			ctx.Id = *res.Data
			ctx.Logger = Logger.With(zap.String("id", ctx.Id))
		}

		// TODO: Fix
		if len(roles) > 0 {
			for _, role := range roles {
				if role == author {
					ctx.User = CtxUser{
						UserId: author,
						Role:   role,
					}
					break
				}
			}
		} else {
			ctx.User = CtxUser{
				UserId: "",
				Role:   "public",
			}
		}

		if ctx.User.Role == "" {
			ctx.SendResponse(ctx.Unauthorized())
			return
		}

		ctx.Logger.Info(url)

		if res := ctx.ParseRequest(); IsError(res) {
			ctx.SendResponse(ctx.BadRequest())
			return
		} else {
			ctx.Req = res.Data
		}

		ctx.SendResponse(handler(ctx))
	})
}

func StartServer(port string) {
	Logger.Info("Starting server on port " + port)

	http.ListenAndServe(":"+port, Server)
}
