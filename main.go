package main

import (
	"go-web/gw"
	"log"
	"net/http"
)

func main() {
	log.Println("go web framework")

	r := gw.New()

	r.Get("/", func(c *gw.Context) {
		c.HTML(http.StatusOK, "<h1>hello world</h1>")
	})

	r.Get("/hello", func(c *gw.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s \n", c.Query("name"), c.Path)
	})

	r.Post("/registry", func(c *gw.Context) {
		c.JSON(http.StatusOK, gw.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	_ = r.Run(":9999")
}
