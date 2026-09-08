package main
import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		fmt.Println("任务A")
	}()

	go func(){
		defer wg.Done()
	fmt.Println("任务B")
	}()
	
	wg.Wait()
	fmt.Println("所有任务完成")
}