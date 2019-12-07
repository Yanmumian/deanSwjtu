package dao

import (
	"deanSwjtu/model"
	"deanSwjtu/utils"
	"fmt"
)


//AddStudent 向数据库中添加注册的学生信息
func AddStudent(student *model.Student) bool {
	sqlStr := "insert into students (stu_name,stu_number,`password`,sex,email,phone,address," +
		"class,department) values (?,?,?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, student.Name, student.Number,student.Password,student.Sex,
		student.Email,student.Phone,student.Address,student.Class,student.Department)
	if err != nil {
		fmt.Println("向数据库插入数据错误,err:",err)
		return false
	}
	return true
}

//QueryStudent 查询该学生是否存在于数据库中
func QueryStudent(stu_number string) bool {
	strSql := "select stu_name from students where stu_number = ?"
	row := utils.Db.QueryRow(strSql, stu_number)
	name := ""
	row.Scan(&name)
	if name == "" {
		return false
	}
	return  true
}

//QueryStudent 查询该学生是否存在于数据库中
func CheckStudent(username ,password string) bool {
	strSql := "select stu_name from students where stu_number = ? and password = ?"
	row := utils.Db.QueryRow(strSql, username, password)
	name := ""
	row.Scan(&name)
	if name == "" {
		return false
	}
	return  true
}