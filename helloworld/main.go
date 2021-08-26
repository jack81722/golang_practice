package main

import (
	"fmt"
	"hello/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	// r.POST("/h", Hello)
	// r.NoRoute(func(ctx *gin.Context) {
	// 	ctx.JSON(404, "See u")
	// })
	// r.Run()
	// fmt.Println("helloworld")
	var p *model.TriplePlus
	// p = &model.Plus{1, 1}
	// fmt.Println(p.Cal())

	p = &model.TriplePlus{
		model.Plus{1, 1},
		1,
	}
	fmt.Println(p.Cal())

}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HelloWorld")
}
