package ajax

import (
	"net/http"
	"encoding/json"
	"fmt"
	"crypto/sha1"
	"go.admin.youqu.com/config"
	"strings"
	"go.admin.youqu.com/application/model"
)

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := strings.TrimSpace(r.PostFormValue("username"))
	password := strings.TrimSpace(r.PostFormValue("password"))
	//加盐值 再sha1
	s := sha1.New()
	s.Write([]byte(password + config.Salt))
	password = fmt.Sprintf("%x", s.Sum(nil))
	name := model.CheckLogin(userName, password)
	info := make(map[string]interface{})
	if(name == ""){
		info["status"] = 1
		info["errorMsg"] = "用户名或密码错误"
	}else{
		info["status"] = 0
		info["errorMsg"] = "成功"
	}
	result, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(result)
}
