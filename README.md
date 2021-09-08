# go web framework
  # Installation
To install Gin package, you need to install Go and set your Go workspace first.

The first need Go installed (version 1.13+ is required), then you can use the below Go command to install Gin.
$ go get -u github.com/gin-gonic/gin
Import it in your code:
import "github.com/gin-gonic/gin"
(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.
import "net/http"
  #Quick start
# assume the following codes in example.go file
$ cat example.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
# run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser
$ go run example.go
