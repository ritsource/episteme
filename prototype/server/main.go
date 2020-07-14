package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/repo"
)

func main() {
	err := repo.Inititalize(".data/dst/prod")
	if err != nil {
		panic(err)
	}

	// posts := repo.GetPostsByCategory(models.Post_Category{Title: "Learning"})
	// fmt.Printf("\"Learning\" Posts ...\n%+v\n", posts)

	r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.GET("/", renderers.RootHandler)

	r.Run()
}
