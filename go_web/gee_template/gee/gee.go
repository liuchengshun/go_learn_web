package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	// 创建一个路由map key => Handle
	router map[string]HandlerFunc
}

func New() *Engine {
	// New() 代表初始化生成一个新的可以使用的map函数
	return &Engine{
		router : make(map[string]HandlerFunc),
	}
}

// 增加路由
func (engine *Engine) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	// 开启端口服务器
	// 因为engine 实现了ServeHTTP 所以将这个函数作为参数传递进去，
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s \n", r.URL)
	}
}

// 上面的总体思想是：使用Method+ routerPath构成key ,再将handlerFunc 作为value
// 当时