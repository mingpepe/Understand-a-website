package webframework

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router      *router
	middlewares []HandlerFunc
}

func New() *Engine {
	return &Engine{
		router:      newRouter(),
		middlewares: make([]HandlerFunc, 0),
	}
}

func (e *Engine) Use(handler HandlerFunc) {
	e.middlewares = append(e.middlewares, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.router.add(http.MethodGet, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.router.add(http.MethodPost, pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r, e.middlewares)
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
