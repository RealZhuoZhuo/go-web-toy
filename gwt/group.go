package gwt

type RouterGroup struct {
	perfix      string
	middlewares []Handlerfunc
	parent      *RouterGroup
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		perfix: group.perfix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handler Handlerfunc) {
	path := group.perfix + "/" + comp
	group.engine.router.addRoute(method, path, handler)
}
func (group *RouterGroup) Get(path string, handler Handlerfunc) {
	group.addRoute("GET", path, handler)
}
func (group *RouterGroup) Post(path string, handler Handlerfunc) {
	group.addRoute("POST", path, handler)
}
func (group *RouterGroup) Use(middlewares ...Handlerfunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
