package main

import (
	"fmt"

	"github.com/ritsource/episteme/prototype/data/models"
	_ "github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/repo"
)

func main() {
	repo.PopulateWithDummyData(models.Posts{
		&models.Post{
			Categories: []*models.Post_Category{
				&models.Post_Category{
					Title: "Physics",
				},
			},
		},
		&models.Post{
			Categories: []*models.Post_Category{
				&models.Post_Category{
					Title: "Computer Science",
				},
				&models.Post_Category{
					Title: "Mathematics",
				},
			},
		},
		&models.Post{
			Categories: []*models.Post_Category{
				&models.Post_Category{
					Title: "Startups",
				},
			},
		},
		&models.Post{
			Categories: []*models.Post_Category{
				&models.Post_Category{
					Title: "Design",
				},
			},
		},
	})

	posts := repo.GetPostsByCategory(models.Post_Category{Title: "Startups"})

	fmt.Printf("\"Startups\" Posts ...\n%+v\n", posts)
}
