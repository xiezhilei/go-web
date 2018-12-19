
package main

import (
	"net/http"
	"fmt"
	"runtime"
	"reflect"
)

//串联多个处理器函数
func hello(w http.ResponseWriter , r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called -" + name)
		h(w,r)
	}
}

func protect(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is protect func")
		h(w,r)
	}
}

//串联多个处理器
type WorldHandler struct {
}
func (h WorldHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"WorldHandler")
}

func logHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("logHandler called - %T\n", h)
			h.ServeHTTP(w,r)
		})
}

func protectHandler(h http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("protectHandler called - %T\n", h)
			h.ServeHTTP(w,r)
		})
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/nihao", protect(log(hello)))

	worldHandler := WorldHandler{}
	http.Handle("/world/", protectHandler(logHandler(worldHandler)))

	server.ListenAndServeTLS("cert.pem", "key.pem")
}
