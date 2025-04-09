package main

import (
	"fmt"
	"gwt"
	"time"
)

func test() gwt.Handlerfunc {
	return func(c *gwt.Context) {
		fmt.Printf("test....")
		t1 := time.Now()
		c.Next()
		t2 := time.Now()
		fmt.Printf("%v", t2.Sub(t1))
	}
}
func main() {
	g := gwt.New()
	g.Use(test())
	g.Get("/test", func(c *gwt.Context) {
		c.W.Write([]byte("hello"))
		time.Sleep(3 * time.Second)
	})
	g.Run(":8888")
}
