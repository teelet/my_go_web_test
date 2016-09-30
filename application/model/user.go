package model

import (
	"database/sql"
	"fmt"
	"strings"
	"go.admin.youqu.com/config"
	_ "github.com/go-sql-driver/mysql"
	"crypto/sha1"
)

type Root struct {
	UserName string
	Auth     string
	Executor string
	Atime    string
}

type RootLog struct {
	UserName string
	Atime    string
	Content  string
	Ip       string
}

var db *sql.DB
var err error
func init() {
	db, err = sql.Open("mysql", "test:123qwe@tcp(139.129.36.196:3306)/gameinfo?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.Ping()
}

func GetRootAuth(userName string) string {
	sql := "select auth from root where username = ?"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	row := stmt.QueryRow(userName)
	var auth string
	row.Scan(&auth)
	return auth
}

func UpdateRoot(userName, auth string)string{
	s := sha1.New()
	s.Write([]byte("123qwe"))
	initPasswd := fmt.Sprintf("%x", s.Sum(nil))
	s = sha1.New()
	s.Write([]byte(initPasswd + config.Salt))
	initPasswd = fmt.Sprintf("%x", s.Sum(nil))

	sql := "insert into root (username, password, auth, excutor) values (?, ?, ?, ?) on duplicate key update auth = ? "
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(userName, initPasswd, auth, "root", auth)
	if err != nil {
		fmt.Println(err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if row >= 1 {
		return "操作成功!"
	}else{
		return "操作失败!"
	}
}

func UpdatePasswd(userName, password string)string{
	sql := "update root set password = ? where username = ?"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(password, userName)
	if err != nil {
		fmt.Println(err)
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if row >= 1 {
		return "操作成功!"
	}else{
		return "操作失败!"
	}
}

func GetLogByPage(page, pageSize int) []RootLog {
	start := (page - 1) * pageSize
	sql := "select username, atime, content, ip from root_log order by atime desc limit ?, ?"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query(start, pageSize)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var logs []RootLog
	for rows.Next() {
		log := RootLog{}
		rows.Scan(&(log.UserName), &(log.Atime), &(log.Content), &(log.Ip))
		logs = append(logs, log)
	}
	return logs
}

func GetLogCount() int {
	sql := "select count(*) from root_log"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}

	return count
}

func CheckLogin(userName, password string)string{
	sql := "select username from root where username = ? and password = ? limit 1"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	row := stmt.QueryRow(userName, password)
	var name string
	row.Scan(&name)
	return name
}

func GetRoots() []Root {
	sql := "select username, auth, excutor, atime from root"
	stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query()
	rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var result []Root

	for rows.Next() {
		root := Root{}
		rows.Scan(&(root.UserName), &(root.Auth), &(root.Executor), &(root.Atime))
		if root.Auth != "*" && root.Auth != "" {
			var newAuthArr []string
			authIds := strings.Split(root.Auth, ",")
			for _, v := range authIds {
				item := config.Auth.Auth[v].AuthName
				newAuthArr = append(newAuthArr, item)
			}
			root.Auth = strings.Join(newAuthArr, ",")
		}
		result = append(result, root)
	}

	return result
}
