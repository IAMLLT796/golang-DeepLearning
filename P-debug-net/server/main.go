package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册对u应请求路径 /ping 的 handler 函数
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req: pong")
		w.Write([]byte("pong"))
	})
	// 调用函数, 启动一个端口为 8080 的 http 服务
	http.ListenAndServe(":8080", nil)
}
