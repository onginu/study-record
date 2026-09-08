package main

import "fmt"

func main(){
ch:=make(chan int)

go func(){
	var sum int=0
	for i:=1;i<=100;i++{
	sum+=i
	
	}
	ch<-sum
}()



num:=<-ch
fmt.Println("结果：",num)

}