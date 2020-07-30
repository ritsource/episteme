package renderers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"text/template"

	"github.com/ritsource/episteme/prototype/constants"
	"github.com/ritsource/episteme/prototype/data/models"
	"github.com/ritsource/episteme/prototype/server/repo"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != RoutesMap.Categories {
		fmt.Printf("r.URL.Path = %+v\n", r.URL.Path)
		NotFoundHandler(w, r)
		return
	}

	categories := repo.GetAllCategories()

	tfd, err := ioutil.ReadFile(path.Join(constants.RepositoryRoot, "prototype/server/static/pages/root.html"))
	if err != nil {
		writeErr(w, 500, errors.New("unable to read template file"))
	}

	t, err := template.New("categoriesPageTemplate").Funcs(
		template.FuncMap{
			"ToLower": strings.ToLower,
		},
	).Parse(string(tfd))
	if err != nil {
		writeErr(w, 500, err)
	}

	err = t.Execute(w, struct {
		Posts                          models.Posts
		Categories                     []models.Post_Category
		SelectedCategoryTitle          string
		SelectedCategoryTitleFormatted string
		RoutesMap                      RoutesMapType
		PageInfo                       struct {
			Page string
		}
	}{
		Posts: models.Posts{
			&models.Post{
				Title: "All Categories",
				Links: func(ctgs []models.Post_Category) []*models.Post_Link {
					res := []*models.Post_Link{}
					for _, ctg := range ctgs {
						// fmt.Printf("- %v\n", fmt.Sprintf("/?category=%v", ctg.GetTitle()))
						res = append(res, &models.Post_Link{
							Src: &models.Post_Link_Src{
								Text: ctg.GetTitle(),
								Url:  fmt.Sprintf("/?category=%v", ctg.GetTitle()),
							},
						})
					}
					return res
				}(categories),
			},
		},
		Categories:                     categories,
		SelectedCategoryTitle:          "",
		SelectedCategoryTitleFormatted: "Pinned Websites",
		RoutesMap:                      RoutesMap,
		PageInfo:                       struct{ Page string }{RoutesMap.Categories},
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
