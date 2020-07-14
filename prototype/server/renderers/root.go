package renderers

import (
	"github.com/gin-gonic/gin"
	"github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/repo"
)

const DefaultCategory = "Learning"

func RootHandler(c *gin.Context) {
	qstrs := c.Request.URL.Query()
	ctgs := qstrs["category"]

	ctg := DefaultCategory
	if len(ctgs) > 0 && ctgs[0] != "" {
		ctg = ctgs[0]
	}

	posts := repo.GetPostsByCategory(models.Post_Category{
		Title: ctg,
	})

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "success",
		"posts":   posts,
	})
}
