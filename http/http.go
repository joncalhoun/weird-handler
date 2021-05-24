package http

import "net/http"

type ResponseWriter = http.ResponseWriter
type Request = http.Request

var (
	ListenAndServe = http.ListenAndServe
)

// HandleFunc replaces the std lib HandleFunc func call
func HandleFunc(pattern string, h HandlerFunc) {
	http.Handle(pattern, toStdLibHandler(h))
}

func toStdLibHandler(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responder := h.Handle(r)
		responder.Respond(w)
	})
}

// Handler replaces http.Handler and doesn't have access to the response writer.
// A responder must be returned to write a response.
type Handler interface {
	Handle(r *Request) Responder
}

type HandlerFunc func(r *Request) Responder

func (hf HandlerFunc) Handle(r *Request) Responder {
	return hf(r)
}

// Responder is used to respond to an http request
type Responder interface {
	Respond(w http.ResponseWriter)
}

type ResponderFunc func(w http.ResponseWriter)

func (rf ResponderFunc) Respond(w http.ResponseWriter) {
	rf(w)
}

// Error is a replacement for the std lib's http.Error
// using the Handler/Responder setup
func Error(err string, code int) ResponderFunc {
	return func(w http.ResponseWriter) {
		http.Error(w, err, code)
	}
}

func OkResponder() ResponderFunc {
	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusOK)
	}
}
