package controller

import (
	"net/http"
	"go.admin.youqu.com/config"
	"fmt"
	"html/template"
	"strings"
	"go.admin.youqu.com/application/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})
	var userName, status string
	if r.Method == "POST" {
		r.ParseForm()
		userName = strings.TrimSpace(r.PostFormValue("username"))
		if userName == "" {
			http.Redirect(w, r, "/login" ,302)
		}
		http.SetCookie(w, &http.Cookie{
			Name  : "user",
			Value : userName,
		})
		http.SetCookie(w, &http.Cookie{
			Name  : "status",
			Value : "login",
		})
	}else{
		cookie, err := r.Cookie("user")
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		userName = cookie.Value
		if userName == "" {
			http.Redirect(w, r, "/login", 302)
			return
		}
		cookie, err = r.Cookie("status")
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		status = cookie.Value
		if status == "logout" {
			http.Redirect(w, r, "/login", 302)
			return
		}
	}

	data["userName"] = userName
	data["status"]   = status

	// 获取权限配置
	class := config.Auth.Class
	auth  := config.Auth.Auth
	userAuth := model.GetRootAuth(userName)

	type ClassItem struct {
		ClassId string
		ClassName string
		AuthList map[string]config.AuthItem
		AuthCount int
	}
	var classList = make(map[string]ClassItem)

	for k, v := range class {
		tmp := ClassItem{ClassId:k, ClassName:v, AuthList:make(map[string]config.AuthItem)}
		for k1, v1 := range auth {
			if v1.ClassId == k{
				if userAuth == "*" {
					tmp.AuthList[k1] = v1
				}else{
					if strings.Index(userAuth, k1) != -1 {
						tmp.AuthList[k1] = v1
					}
				}
			}
		}
		tmp.AuthCount = len(tmp.AuthList)

		classList[k] = tmp
	}
	data["auth"] = classList

	htmlStr, err := config.GetHtmlByPath("application/view/main.html")
	if err != nil {
		fmt.Print(err)
	}
	tpl, err := template.New("main").Parse(string(htmlStr))
	if err != nil {
		fmt.Print(err)
	}
	tpl.Execute(w, data)
}
