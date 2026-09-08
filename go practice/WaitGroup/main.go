package main
import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(5)
	for i:=1;i<=5;i++{
		go func(n int){
		defer wg.Done()
		fmt.Println("任务",n,"完成")
		}(i)
		
	}
	wg.Wait()
	fmt.Println("全部完成")
}