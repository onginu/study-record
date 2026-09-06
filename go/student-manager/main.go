package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	students := []Student{}
	for {
		fmt.Println("学生管理系统 ")
		fmt.Println("1.添加")
		fmt.Println("2.删除")
		fmt.Println("3.修改")
		fmt.Println("4.查询")
		fmt.Println("5.显示全部")
		fmt.Println("0.退出")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:

			students = addStudent(students)
		case 2:

			students = deleteStudent(students)
		case 3:

			students = changeStudent(students)
		case 4:

			findStudent(students)
		case 5:

			showstudent(students)
		case 0:

			return

		}
	}
}
func addStudent(students []Student) []Student {
	var id int
	var name string
	var score float64
	fmt.Println("请输入添加的学生的id：")
	fmt.Scan(&id)
	fmt.Println("请输入添加的学生的name：")
	fmt.Scan(&name)
	fmt.Println("请输入添加的学生的score：")
	fmt.Scan(&score)
	student := Student{
		ID:    id,
		Name:  name,
		Score: score,
	}
	students = append(students, student)
	fmt.Println("添加成功！")
	return students
}

func showstudent(students []Student) {
	if len(students) == 0 {
		fmt.Println("暂无学生")
		return
	}

	for _, v := range students {
		fmt.Println("该学生的id为：")
		fmt.Println(v.ID)
		fmt.Println("该学生的name为：")
		fmt.Println(v.Name)
		fmt.Println("该学生的score为：")
		fmt.Println(v.Score)
	}
}

func findStudent(students []Student) {
	var _id int
	fmt.Println("输入要查找的学生id：")
	fmt.Scan(&_id)
	var found bool = false
	for _, v := range students {

		if v.ID == _id {
			found = true
			fmt.Println("该学生的id为：")
			fmt.Println(v.ID)
			fmt.Println("该学生的name为：")
			fmt.Println(v.Name)
			fmt.Println("该学生的score为：")
			fmt.Println(v.Score)
			return
		}

	}
	if found == false {
		fmt.Println("查无此人")
	}
}

func deleteStudent(students []Student) []Student {
	var _id int
	fmt.Println("输入要查找的学生id：")
	fmt.Scan(&_id)

	for i, v := range students {
		if v.ID == _id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("删除成功")
			return students
		}
	}

	fmt.Println("查无此人")
	return students
}

func changeStudent(students []Student) []Student {
	var _id int
	fmt.Println("输入要修改的学生id：")
	fmt.Scan(&_id)

	for i := range students {
		if students[i].ID == _id {
			var _name string
			var _score float64
			fmt.Println("输入要修改该学生的name为：")
			fmt.Scan(&_name)
			fmt.Println("输入要修改该学生的score为：")
			fmt.Scan(&_score)
			students[i].Name = _name
			students[i].Score = _score
			fmt.Println("修改成功！")
			return students
		}
	}
	fmt.Println("查无此人")
	return students
}
