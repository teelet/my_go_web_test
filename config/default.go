package config

import (
	"io/ioutil"
	"encoding/json"
)

var (
	Salt = "youqu_salt"  //密码盐值
	needCacheHtml = false //是否使用html缓存,默认true(因为golang的template渲染html时,每次都重新读取html文件)
	HtmlCache = make(map[string]string) //存放html缓存
	Auth AllAuth //权限配置
)

type AuthItem struct {
	AuthName string `json:"authName"`
	ClassId  string `json:"classId"`
	Url      string `json:"url"`
}

type classList map[string]string
type authList map[string]AuthItem

type AllAuth struct {
	Class classList `json:"class"`
	Auth  authList  `json:"auth"`
}

func init() {
	jsonByte, _ := ioutil.ReadFile("config/auth.json")
	json.Unmarshal(jsonByte, &Auth)
}

func GetHtmlByPath(htmlFilePath string) (string, error) {
	if needCacheHtml {
		if str, ok := HtmlCache[htmlFilePath]; ok {
			return str, nil
		}else{
			htmlByte, err := ioutil.ReadFile(htmlFilePath)
			str := string(htmlByte)
			if err != nil {
				return "", err
			}else{
				HtmlCache[htmlFilePath] = str
				return str, nil
			}

		}
	}else{
		htmlByte, err := ioutil.ReadFile(htmlFilePath)
		str := string(htmlByte)
		if err != nil {
			return "", err
		}else{
			return str, nil
		}
	}

}