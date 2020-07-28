package repo

import (
	"sort"
	"strings"
	"sync"

	"reflect"

	"github.com/ritsource/episteme/prototype/constants"
	"github.com/ritsource/episteme/prototype/data/models"
)

var Data = struct {
	Posts          models.Posts
	Categories     []models.Post_Category
	PinnedWebsites models.PinnedWebsites
}{
	Posts:          models.Posts{},
	Categories:     []models.Post_Category{},
	PinnedWebsites: models.PinnedWebsites{},
}

func Inititalize() error {
	_, err := Data.Posts.ReadFromFS(constants.DEFAULT_POSTS_DATA_OUTPUT_FILEPATH)
	if err != nil {
		return err
	}

	ctgmap := map[string]bool{}
	for _, post := range Data.Posts {
		for _, ctg := range post.GetCategories() {
			ctgmap[ctg.Title] = true
		}
	}

	for _, title := range reflect.ValueOf(ctgmap).MapKeys() {
		Data.Categories = append(Data.Categories, models.Post_Category{
			Title: title.String(),
		})
	}

	sort.Slice(Data.Categories, func(i, j int) bool {
		return Data.Categories[i].Title < Data.Categories[j].Title
	})

	_, err = Data.PinnedWebsites.ReadFromFS(constants.DEFAULT_PINNED_WEBSITES_DATA_OUTPUT_FILEPATH)
	if err != nil {
		return err
	}

	return nil
}

var CategoryIndex = map[string]models.Posts{}

func GetPostsByCategory(ctg models.Post_Category) (string, models.Posts) {
	// using memoization
	if val, exists := CategoryIndex[strings.ToLower(ctg.GetTitle())]; exists {
		return ctg.GetTitle(), val
	}

	ctgTitle := ""

	posts := make(models.Posts, len(Data.Posts))

	wg := sync.WaitGroup{}

	for idx, post := range Data.Posts {
		wg.Add(1)

		go func(p *models.Post, i int, wg *sync.WaitGroup) {
			for _, category := range p.GetCategories() {
				// NOTE: replace with proto.Equal

				if strings.ToLower(category.GetTitle()) == strings.ToLower(ctg.GetTitle()) {
					posts[i] = p
					ctgTitle = category.Title
					break
				}
			}

			wg.Done()
		}(post, idx, &wg)
	}

	wg.Wait()

	ps := models.Posts{}
	for _, p := range posts {
		if p != nil {
			ps = append(ps, p)
		}
	}

	// using memoization
	CategoryIndex[strings.ToLower(ctgTitle)] = ps

	return ctgTitle, ps
}

func GetAllCategories() []models.Post_Category {
	return Data.Categories
}

func GetPinnedWebsites() models.PinnedWebsites {
	return Data.PinnedWebsites
}
