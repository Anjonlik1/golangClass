package main 

func generator()<- chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch

}
func doubler(in <-chan int) {
	ch := make(chan int)
	for i := range in {
		ch <- i * 2
	}
	close(ch)
}()
return ch

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func main() { 	
	in := generator()
	out := doubler(in)
	printer(out)
	fmt.Println("All done")
	fmt.Println("Exiting main")

}	