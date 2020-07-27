package main

import (
	"fmt"
	"net/http"
	"path"

	_ "github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/constants"
	"github.com/ritsource/episteme/prototype/server/renderers"
	"github.com/ritsource/episteme/prototype/server/repo"
	"github.com/rs/cors"
)

func main() {
	err := repo.Inititalize(".data/dst/prod")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("Data.Categories = %+v\n", repo.Data.Categories)

	// posts := repo.GetPostsByCategory(models.Post_Category{Title: "Learning"})
	// fmt.Printf("\"Learning\" Posts ...\n%+v\n", posts)

	mux := http.NewServeMux()

	staticDir := path.Join(constants.RepositoryRoot, "prototype/server/static")
	fs := http.FileServer(http.Dir(staticDir))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", renderers.RootHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)

	fmt.Printf("Server running on port %+v\n", 8080)
	http.ListenAndServe(":8080", handler)

}
