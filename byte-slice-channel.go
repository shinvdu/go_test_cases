package main
import "fmt"
import "sync"


func main() {
    // fmt.Println("hello world")
    var wg sync.WaitGroup

	tch := make(chan []byte)
    // defer close(tch)

	wg.Add(1)
	go process(&wg, tch)

	for v := range tch {
		fmt.Printf("The value is %s\n", v)
	}

}

func process(wg *sync.WaitGroup, tch chan []byte) {  

    fmt.Println("started Goroutine ")
	tch <- []byte("hello")	
	tch <- []byte("wrold")	
	tch <- []byte("gre")	
	tch <- []byte("start")	
	tch <- []byte("tianbao ")	
	tch <- []byte("ababa")	

    // time.Sleep(2 * time.Second)
	close(tch)
    fmt.Println("Goroutine ended")
    wg.Done()
}
