package gwt

import (
	"net/http"
)

type Handlerfunc func(c *Context)
type Engine struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func (engine *Engine) Get(path string, handler Handlerfunc) {
	engine.router.addRoute("GET", path, handler)
}
func (engine *Engine) Post(path string, handler Handlerfunc) {
	engine.router.addRoute("POST", path, handler)
}
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	engine.router.Handle(context)
}
