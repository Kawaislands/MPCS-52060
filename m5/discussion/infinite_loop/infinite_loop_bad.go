package main
import "fmt"
import "time"
import "runtime"

func main() {
    var result int
	processors := runtime.GOMAXPROCS(-1) 
	fmt.Println(processors)
    for i := 0; i < processors; i++ {
        go func() {
            for { result++ }
        }()
    }
    time.Sleep(3*time.Second)       //wait for go function to increment the value.
    fmt.Println("result =", result)
}