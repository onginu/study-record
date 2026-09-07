package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"io"
)

type StudentManager struct{
	students []Student
	
}

func (m *StudentManager)Save(){
	data,err:=json.Marshal(m.students)
	if err!=nil{
    fmt.Println(err)
    return

}
	file,err:=os.OpenFile(
		"students.json",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0666,

	)
	if err!=nil{
    fmt.Println(err)
    return

}
defer file.Close()


	_, err = file.Write(data)
if err != nil {
	fmt.Println(err)
	return
}


}

func (m *StudentManager)Load(){
	
file,err:=os.Open("students.json")
if err!=nil{
	if os.IsNotExist(err){
		return
	}
	
    fmt.Println(err)
    return

}

defer file.Close()

data,err:=io.ReadAll(file)
if err!=nil{
    fmt.Println(err)
    return

}


	err=json.Unmarshal(data,&m.students)
	if err!=nil{
		fmt.Println(err)
		return 
	}
	
}

func (m*StudentManager)AddStudent() {
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
	m.students = append(m.students, student)
	m.Save()
	fmt.Println("添加成功！")
	
}

func (m*StudentManager)Showstudent() {
	if len(m.students) == 0 {
		fmt.Println("暂无学生")
		return
	}

	for _, v := range m.students {
		fmt.Println("该学生的id为：")
		fmt.Println(v.ID)
		fmt.Println("该学生的name为：")
		fmt.Println(v.Name)
		fmt.Println("该学生的score为：")
		fmt.Println(v.Score)
	}
}

func (m*StudentManager)FindStudent() {
	var _id int
	fmt.Println("输入要查找的学生id：")
	fmt.Scan(&_id)
	var found bool = false
	for _, v := range m.students {

		if v.ID == _id {
			found = true
			fmt.Println("该学生的id为：")
			fmt.Println(v.ID)
			fmt.Println("该学生的name为：")
			fmt.Println(v.Name)
			fmt.Println("该学生的score为：")
			fmt.Println(v.Score)
			
		}

	}
	if found == false {
		fmt.Println("查无此人")
	}
}

func (m*StudentManager)DeleteStudent()  {
	var _id int
	fmt.Println("输入要查找的学生id：")
	fmt.Scan(&_id)

	for i, v := range m.students {
		if v.ID == _id {
			m.students = append(m.students[:i], m.students[i+1:]...)
			m.Save()
			fmt.Println("删除成功")

			return 
		}
	}

	fmt.Println("查无此人")
	
}

func (m*StudentManager)ChangeStudent() {
	var _id int
	fmt.Println("输入要修改的学生id：")
	fmt.Scan(&_id)

	for i := range m.students {
		if m.students[i].ID == _id {
			var _name string
			var _score float64
			fmt.Println("输入要修改该学生的name为：")
			fmt.Scan(&_name)
			fmt.Println("输入要修改该学生的score为：")
			fmt.Scan(&_score)
			m.students[i].Name = _name
			m.students[i].Score = _score
			m.Save()
			fmt.Println("修改成功！")
			return 
		}
	}
	fmt.Println("查无此人")

}

func (m *StudentManager) SortByScore(){
	sort.Slice(m.students,func(i,j int)bool{
		return m.students[i].Score>m.students[j].Score
	})
	fmt.Println("学生成绩排名为：")
	for _,v:=range m.students{
		
		fmt.Println(v.Name,v.Score)
	}
}

func(m*StudentManager)AverageScore(){

	if len(m.students)==0{
	fmt.Println("暂无学生")
	return
}

sum:=0.0
for _,v:=range m.students{
	sum+=v.Score
}

fmt.Println("班级平均分为：")
n:=sum/float64(len(m.students))
fmt.Println(n)
}

func(m*StudentManager)MaxScore(){
	if len(m.students)==0{
	fmt.Println("暂无学生")
	return
}

	msx:=m.students[0].Score
	for _,v:=range m.students{
	if v.Score>msx{
	msx=v.Score
	}
	}

	fmt.Println("班级最高分为：")
	fmt.Println(msx)
}


func(m*StudentManager)SearchByName(){
	var name string
	fmt.Println("请输入你要查找的姓名：")
	fmt.Scan(&name)
	found:=false
	for _,v:=range m.students{
		if name==v.Name{
			found=true
			fmt.Println("该学生的id为：")
			fmt.Println(v.ID)
			fmt.Println("该学生的分数为：")
			fmt.Println(v.Score)
		}

	}
	if !found{
		fmt.Println("查无此人")
	}
}