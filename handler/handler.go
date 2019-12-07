package handler

import (
	"deanSwjtu/dao"
	"deanSwjtu/model"
	"html/template"
	"io/ioutil"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("./index.html"))
		t.Execute(w,"")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	if dao.CheckStudent(username,password) {
		t := template.Must(template.ParseFiles("./views/login_successful.html"))
		t.Execute(w,"")
		return
	}
	t := template.Must(template.ParseFiles("./views/login.html"))
	t.Execute(w,"")
}

func RegistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("./views/regist.html"))
		t.Execute(w,"")
		return
	}
	name := r.PostFormValue("name")
	number := r.PostFormValue("number")
	password := r.PostFormValue("password")
	sex := r.PostFormValue("sex")
	email := r.PostFormValue("email")
	phone := r.PostFormValue("phone")
	address := r.PostFormValue("address")
	class := r.PostFormValue("email")
	department := r.PostFormValue("department")

	student := &model.Student{
		Name:     name,
		Number:   number,
		Password: password,
		Sex:      sex,
		Email:email,
		Phone:phone,
		Address:address,
		Class:class,
		Department:department,
	}

	//加入该判断是为了不让在刷新时重新提交post，使得相同的学生注册数据插入到数据库，避免出错
	if dao.QueryStudent(number) {
		bytes, _ := ioutil.ReadFile("./views/login_successful.html")
		w.Write(bytes)
		return
	}

	//若数据库中不存在该学生信息，则才能够添加
	dao.AddStudent(student)
	bytes, _ := ioutil.ReadFile("./views/login_successful.html")
	w.Write(bytes)
}
