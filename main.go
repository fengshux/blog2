package main

import (
	"log"

	"github.com/fengshux/blog2/backend"
	"github.com/fengshux/blog2/build"
	"github.com/gin-gonic/gin"
)

func main() {
	// check and init project
	build.InitProject()

	// start server
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Static("/pages", "./pages")

	reg, err := backend.NewRegister()
	if err != nil {
		panic(err)
	}
	reg.Regist(r.Group("/api"))

	r.HEAD("/", func(c *gin.Context) {
		c.Done()
	})
	log.Println("server listen on 8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
