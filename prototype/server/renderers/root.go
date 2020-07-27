package renderers

import (
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
	if ctg == "" {
		ctg = DefaultCategory
	}

	ctgTitle, posts := repo.GetPostsByCategory(models.Post_Category{
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

	ps := models.Posts{}
	for _, p := range posts {
		ps = append(ps, p)
	}
	for _, p := range posts {
		ps = append(ps, p)
	}
	for _, p := range posts {
		ps = append(ps, p)
	}
	for _, p := range posts {
		ps = append(ps, p)
	}

	posts = ps

	// cs := []models.Post_Category{}
	// for i := 0; i < 40; i++ {
	// 	for _, c := range categories {
	// 		cs = append(cs, c)
	// 	}
	// }

	// categories = cs

	// // ctg formatted
	// ctgf := func(str string) string {
	// 	res := ""
	// 	for _, s := range strings.Split(str, " ") {
	// 		if s != "" {
	// 			res += strings.ToUpper(s[0:1]) + strings.ToLower(s[1:len(s)])
	// 			res += " "
	// 		}
	// 	}
	// 	return res
	// }(ctg)

	err = t.Execute(w, struct {
		Posts                 models.Posts
		Categories            []models.Post_Category
		SelectedCategoryTitle string
	}{
		Posts:                 posts,
		Categories:            categories,
		SelectedCategoryTitle: ctgTitle,
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
