package repo

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"reflect"

	"github.com/ritsource/episteme/prototype/data/models"
)

// var Data models.Posts = models.Posts{}

var Data = struct {
	Posts      models.Posts
	Categories []models.Post_Category
}{
	Posts:      models.Posts{},
	Categories: []models.Post_Category{},
}

func PopulateWithDummyData(ps models.Posts) {
	Data.Posts = ps
}

func Inititalize(fp string) error {
	fpAbs, err := filepath.Abs(fp)
	if err != nil {
		fmt.Printf("could not resolve file %v\n", fp)
		return err
	}

	_, err = Data.Posts.ReadFromFS(fpAbs)
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
