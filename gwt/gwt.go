package gwt

import (
	"net/http"
)

type Handlerfunc func(http.ResponseWriter, *http.Request)
type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) Get(pattern string, handler Handlerfunc) {
	engine.router.addRoute("GET", pattern, handler)
}
func (engine *Engine) Post(pattern string, handler Handlerfunc) {
	engine.router.addRoute("POST", pattern, handler)
}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router.handlers[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}
