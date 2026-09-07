package main

import (
"fmt"

)


func main() {
	manager:=StudentManager{}
	manager.Load()
	for {
		fmt.Println("学生管理系统 ")
		fmt.Println("1.添加")
		fmt.Println("2.删除")
		fmt.Println("3.修改")
		fmt.Println("4.查询")
		fmt.Println("5.显示全部")
		fmt.Println("6.按成绩排序")
		fmt.Println("7.平均成绩")
		fmt.Println("8.最高分")
		fmt.Println("9.按姓名搜索")
		fmt.Println("0.退出")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:

			manager.AddStudent()
			manager.Save()
		case 2:

			manager.DeleteStudent()
		case 3:

			manager.ChangeStudent()
		case 4:

			manager.FindStudent()
		case 5:

			manager.Showstudent()
		case 6:
			manager.SortByScore()
		case 7:
			manager.AverageScore()
		case 8:
			manager.MaxScore()
		case 9:
			manager.SearchByName()		
		case 0:
			manager.Save()
			return

		}
	}
}

