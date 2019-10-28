package main

import (
	"fmt"
	"github.com/qingwenjie/gpool"
	"time"
)

func main() {
	gp := gpool.New(2)
	for i := 1; i <= 100; i++ {
		gp.Add(1)
		go func(j int) {
			defer gp.Done()
			fmt.Println(j)
			time.Sleep(1 * time.Second)
		}(i)
	}
	gp.Wait()
}
