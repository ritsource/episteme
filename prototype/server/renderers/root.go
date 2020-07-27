package renderers

import (
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/constants"
	"github.com/ritsource/episteme/prototype/server/repo"
)

const DefaultCategory = "Learning"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	qstrs := r.URL.Query()

	ctg := string(qstrs.Get("category"))
	fmt.Printf("ctg = %+v\n", ctg)
	if ctg == "" {
		ctg = DefaultCategory
	}

	posts := repo.GetPostsByCategory(models.Post_Category{
		Title: ctg,
	})
	categories := repo.GetAllCategories()

	if r.URL.Path != "/" {
		NotFoundHandler(w, r)
		return
	}

	t, err := template.ParseFiles(
		path.Join(constants.RepositoryRoot, "prototype/server/static/pages/root.html"),
	)
	if err != nil {
		writeErr(w, 500, err)
	}

	fmt.Printf("posts = %+v\n", posts)

	err = t.Execute(w, struct {
		Posts      models.Posts
		Categories []models.Post_Category
	}{
		Posts:      posts,
		Categories: categories,
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
