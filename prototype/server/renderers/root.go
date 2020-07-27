package renderers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"text/template"

	"github.com/ritsource/episteme/prototype/constants"
	"github.com/ritsource/episteme/prototype/data/models"
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

	tfd, err := ioutil.ReadFile(path.Join(constants.RepositoryRoot, "prototype/server/static/pages/root.html"))
	if err != nil {
		writeErr(w, 500, errors.New("unable to read template file"))
	}

	t, err := template.New("rootPageTemplate").Funcs(
		template.FuncMap{
			"ToLower": strings.ToLower,
		},
	).Parse(string(tfd))

	// t, err := template.ParseFiles(
	// 	path.Join(constants.RepositoryRoot, "prototype/server/static/pages/root.html"),
	// )

	if err != nil {
		writeErr(w, 500, err)
	}

	err = t.Execute(w, struct {
		Posts                          models.Posts
		Categories                     []models.Post_Category
		SelectedCategoryTitle          string
		SelectedCategoryTitleFormatted string
	}{
		Posts:                 posts,
		Categories:            categories,
		SelectedCategoryTitle: ctgTitle,
		SelectedCategoryTitleFormatted: func(str string) string {
			// formatting title
			res := ""
			for _, s := range strings.Split(str, " ") {
				if s != "" {
					res += strings.ToUpper(s[0:1]) + strings.ToLower(s[1:len(s)])
					res += " "
				}
			}
			return strings.TrimSpace(res)
		}(ctgTitle),
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
