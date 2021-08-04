package main
import (
"golang.org/x/tour"
"fmt"
)

func Walc (t *tree.Tree, ch chan int) {
	recWalc(t, ch)
	close(ch)
}

func recWalc(t *tree.Tree, ch chan int) {
	if t !=nil {
		recWalc(t.Left, ch)
		ch<-t.Value
		recWalc(t.Right, ch)
	}
}
func Same(t1,t2 *tree.Tree) bool {
	ch1 := make (chan int)
	ch2 := make (chan int)
	go Walc(t1,ch1)
	go Walc(t1,ch2)

	for {
		x1,ok1 :=<-ch1
		x2,ok2 :=<-ch2
		switch {
		case ok1 != ok2:
			return false
		case !ok1:
			return true
		case x1 !=x2:
			return false
			default
			
		}
	}
}

func main() {
	ch := make(chan int)
	go Walc(tree.New(1),ch)
	for v := range ch{
		fmt.Println(v)
	}
	fmt.Println(Same(tree.New(1),tree.New(1)))
	fmt.Println(Same(tree.New(1),tree.New(2)))
}
