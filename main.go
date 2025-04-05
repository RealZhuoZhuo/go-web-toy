package main

import (
	"gwt"
)

func main() {
	g := gwt.New()
	v1 := g.Group("/v1")
	{
		v1.Get("api", func(c *gwt.Context) {
			c.W.Write([]byte("/v1/api....."))
		})
	}
	v2 := g.Group("/v2")
	{
		v2.Get("api", func(c *gwt.Context) {
			c.W.Write([]byte("/v2/api....."))
		})
	}
	g.Run(":8888")
}
