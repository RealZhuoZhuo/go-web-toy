package main

import (
	"gwt"
)

func main() {
	g := gwt.New()
	g.Get("/test", func(c *gwt.Context) {
		c.W.Write([]byte("test...."))
	})
	g.Run(":8888")
}
