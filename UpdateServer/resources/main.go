package main

import (
	"fmt"
	"net/http"

	"./utillog"
)

func main() {

	defer func() {

		if r := recover(); r != nil {
			utillog.Instance().Fatal(r)
		}
	}()

	// panic(55)

	fmt.Println("启动服务器...")
	utillog.Instance().Init()
	utillog.Instance().Info("启动服务器")
	utillog.Instance().Error("启动服务器")

	router := NewRouter()

	utillog.Instance().Fatal(http.ListenAndServe(":8080", router))
	// utillog.Instance().Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", router))
}
