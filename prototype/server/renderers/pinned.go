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

func PinnedHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/pinned/" {
		fmt.Printf("r.URL.Path = %+v\n", r.URL.Path)
		NotFoundHandler(w, r)
		return
	}

	categories := repo.GetAllCategories()
	pinnedWebsites := repo.GetPinnedWebsites()

	tfd, err := ioutil.ReadFile(path.Join(constants.RepositoryRoot, "prototype/server/static/pages/root.html"))
	if err != nil {
		writeErr(w, 500, errors.New("unable to read template file"))
	}

	t, err := template.New("rootPageTemplate").Funcs(
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
	}{
		Posts: models.Posts{
			&models.Post{
				Title: "More websites like this one",
				Links: func(pws models.PinnedWebsites) []*models.Post_Link {
					res := []*models.Post_Link{}
					for _, w := range pws {
						res = append(res, &models.Post_Link{
							Src: &models.Post_Link_Src{
								Text: w.GetValue(),
								Url:  w.GetValue(),
							},
						})
					}
					return res
				}(pinnedWebsites),
			},
		},
		Categories:                     categories,
		SelectedCategoryTitle:          "",
		SelectedCategoryTitleFormatted: "Pinned Websites",
	})
	if err != nil {
		writeErr(w, 500, err)
	}
}
