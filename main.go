package main

import (
	"deanSwjtu/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)

	//登录，注册，注销
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/regist", handler.RegistHandler)
	http.ListenAndServe(":8080", nil)
}
