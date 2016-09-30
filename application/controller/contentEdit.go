package controller

import (
	"net/http"
	"go.admin.youqu.com/config"
	"fmt"
	"html/template"
)

func ContentEdit(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("user"); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound);
		return
	}
	htmlStr, err := config.GetHtmlByPath("application/view/contentEdit.html")
	if err != nil {
		fmt.Print(err)
	}
	tpl, err := template.New("contentEdit").Parse(string(htmlStr))
	if err != nil {
		fmt.Print(err)
	}
	tpl.Execute(w, nil)
}
