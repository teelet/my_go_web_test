package main

import (
	"fmt"
	"net/http"
	"go.admin.youqu.com/application/controller"
	"go.admin.youqu.com/application/controller/ajax"
	"time"
)


var err error

func main() {
	fmt.Println("server is running...")
	server := http.Server{
		Addr: ":8888",
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//controller
func init() {
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/contentManage", controller.ContentManage)
	http.HandleFunc("/changeKnowledgeDb", controller.ChangeKnowledgeDb)
	http.HandleFunc("/addArticle", controller.AddArticle)
	http.HandleFunc("/addTag", controller.AddTag)
	http.HandleFunc("/addAlias", controller.AddAlias)
	http.HandleFunc("/authManage", controller.AuthManage)
	http.HandleFunc("/authManage/updatePasswd", controller.UpdatePasswd)
	http.HandleFunc("/authManage/updateRoot", controller.UpdateRoot)
	http.HandleFunc("/contentEdit", controller.ContentEdit)
	http.HandleFunc("/image", controller.CheckCode)
}

//ajax
func init() {
	http.HandleFunc("/ajax/checklogin", ajax.CheckLogin)
}

//file server
func init() {
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("static/image/"))))
}
