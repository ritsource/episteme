package repo

import (
	"fmt"
	"path/filepath"
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

	return nil
}

func GetPostsByCategory(ctg models.Post_Category) (string, models.Posts) {
	// NOTE: kind of an hack, so try some other way
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

	return ctgTitle, ps

}

func GetAllCategories() []models.Post_Category {
	return Data.Categories
}
