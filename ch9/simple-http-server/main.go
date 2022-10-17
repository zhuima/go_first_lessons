package main

import (
	"net/http"
)

func main() {
	// r代表来自客户端的HTTP请求，w则是用来操作返回给客户端的应答
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("hello, world")))
	})

	// 建立一个HTTP服务
	http.ListenAndServe(":8080", nil)
}
