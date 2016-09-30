package controller

import (
	"net/http"
	"html/template"
	"fmt"
	"go.admin.youqu.com/config"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userName string = ""
	from := r.URL.Query().Get("from")
	if from == "logout" {
		http.SetCookie(w, &http.Cookie{
			Name    : "status",
			Value   : "logout",
			Expires : time.Now().AddDate(0, 0, 1),
		})
		cookie, _ := r.Cookie("user")
		userName = cookie.Value
	}
	data := map[string]string{
		"userName" : userName,
	}

	w.Header().Add("content-type", "text/html")
	//tpl, err := template.ParseFiles("application/view/login.html");
	htmlStr, err := config.GetHtmlByPath("application/view/login.html")
	if err != nil {
		fmt.Print(err)
	}
	tpl, err := template.New("login").Parse(string(htmlStr))
	if err != nil {
		fmt.Print(err)
	}
	tpl.Execute(w, data)
}