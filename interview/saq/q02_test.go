package saq

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{"zhou", 24},
		{"li", 23},
		{"wang", 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for i, v := range m {
		fmt.Println(i, v.Age, v.Name)
	}
	fmt.Println("--------------")
	// 正确写法
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for i, v := range m {
		fmt.Println(i, v.Age, v.Name)
	}
}

func Test_pase_student(t *testing.T) {
	pase_student()
	complex128
}
