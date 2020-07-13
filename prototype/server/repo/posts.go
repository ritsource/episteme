package repo

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ritsource/episteme/prototype/data/models"
)

var Data models.Posts = models.Posts{}

func PopulateWithDummyData(ps models.Posts) {
	Data = ps
}

func Inititalize(fp string) error {
	fpAbs, err := filepath.Abs(fp)
	if err != nil {
		fmt.Printf("could not resolve file %v\n", fp)
		return err
	}

	_, err = Data.ReadFromFS(fpAbs)
	return err
}

func GetPostsByCategory(ctg models.Post_Category) models.Posts {
	posts := make(models.Posts, len(Data))

	wg := sync.WaitGroup{}

	for idx, post := range Data {
		wg.Add(1)

		go func(p *models.Post, i int, wg *sync.WaitGroup) {
			for _, category := range p.GetCategories() {
				// NOTE: replace with proto.Equal
				if category.GetTitle() == ctg.GetTitle() {
					posts[i] = p
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
	return ps

}
