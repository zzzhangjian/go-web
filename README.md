# go web framework

  # Installation

To install go-web package, you need to install Go and set your Go workspace first.

The first need Go installed (version 1.13+ is required), then you can use the below Go command to install Gin.
`go get -u github.com/zzzhangjian/go-web`
Import it in your code:
`import "github.com/zzzhangjian/go-web"`
(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.
`import "net/http"`

  # Quick start

`# Assume the following codes in example.go file`
`$ cat example.go`

`package main`

`import "github.com/zzzhangjian/go-web"`

`func main() {`
	`r := gw.Default()`
	`r.GET("/ping", func(c *gin.Context) {`
		`c.JSON(200, gw.H{`
			`"message": "pong",`
		`})`
	`})`
	`r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")`
`}`

`# run example.go and visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser`
`$ go run example.go`
