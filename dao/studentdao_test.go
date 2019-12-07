package dao

import (
	"deanSwjtu/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试开始：")
	m.Run()
}

func TestStudent(t *testing.T) {
	fmt.Println("测试student中的函数")
	//t.Run("111",testStudenAdd)
	t.Run("111",testqueryStudent)
}

func testStudenAdd(t *testing.T) {
	student := &model.Student{
		Name:   "111",
		Number: "111",
		Password: "111",
		Sex: "1",
	}
	AddStudent(student)
}

func testqueryStudent(t  *testing.T) {
	b := QueryStudent("111123")
	fmt.Println(b)
}
