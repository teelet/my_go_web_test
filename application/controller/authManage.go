package controller

import (
	"net/http"
	"go.admin.youqu.com/config"
	"fmt"
	"html/template"
	"go.admin.youqu.com/application/model"
	"strconv"
	"strings"
	"crypto/sha1"
)

func UpdateRoot (w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	userName := strings.TrimSpace(r.PostFormValue("username"))
	authes := r.PostForm["auth[]"] //获取checkbox数据
	auth := strings.Join(authes, ",")
	res := model.UpdateRoot(userName, auth)
	r.Form.Set("res_msg", res)
	AuthManage(w, r)
}

func UpdatePasswd(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	userName := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	s := sha1.New()
	s.Write([]byte(password + config.Salt))
	password = fmt.Sprintf("%x", s.Sum(nil))
	res := model.UpdatePasswd(userName, password)
	r.Form.Set("res_msg", res)
	AuthManage(w, r)
}

func AuthManage(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("user"); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound);
		return
	}
	var data = make(map[string]interface{})
	r.ParseForm()
	data["res_msg"]= r.FormValue("res_msg")
	//获取管理员列表
	roots := model.GetRoots()
	data["roots"] = roots

	//获取log日志
	var page, count, pageSize int = 0, 0, 1
	page, _ = strconv.Atoi(r.FormValue("page"))
	if page == 0 {
		page = 1
	}
	count = model.GetLogCount()
	//总页数
	var ceil = 0
	if (count % pageSize) > 0 {
		ceil = 1
	}
	data["allPage"] = count / pageSize + ceil
	data["currentPage"] = page
	//分页提取日志
	data["logs"] = model.GetLogByPage(page, pageSize)

	htmlStr, err := config.GetHtmlByPath("application/view/authManage.html")
	if err != nil {
		fmt.Print(err)
	}
	tpl, err := template.New("authManage").Parse(string(htmlStr))
	if err != nil {
		fmt.Print(err)
	}
	cookie, _ := r.Cookie("user")
	data["user"] = cookie.Value
	data["auth"] = config.Auth.Auth
	tpl.Execute(w, data)
}
