package models_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ritsource/episteme/prototype/data/models"
	"google.golang.org/protobuf/proto"
)

var FILEPATH string
var POSTS models.Posts

func init() {
	fprel := "../../.data/tests/"

	var err error
	FILEPATH, err = filepath.Abs(fprel)
	if err != nil {
		fmt.Printf("could not resolve file path %v\n", fprel)
		panic(err)
	}

	for i := 0; i < 3; i++ {
		POSTS = append(POSTS, &models.Post{})
	}
}

func TestWriteToFS(t *testing.T) {
	err := POSTS.SaveToFS(FILEPATH)
	if err != nil {
		t.Error(err)
	}

	posts := models.Posts{}
	_, err = posts.ReadFromFS(FILEPATH)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("\nPOSTS ...\n%+v\n", POSTS)
	fmt.Printf("\nposts ...\n%+v\n", posts)

	isEq := compareFields(POSTS, posts)
	if !isEq {
		t.Errorf("POSTS and posts are not equal")
	}

}

func compareFields(ps1, ps2 models.Posts) bool {
	if len(ps1) != len(ps2) {
		return false
	}

	for i, p1 := range ps1 {
		p2 := ps2[i]

		if !proto.Equal(p1, p2) {
			return false
		}
	}

	return true
}
