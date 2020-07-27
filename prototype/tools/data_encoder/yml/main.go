package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ritsource/episteme/prototype/constants"
	"github.com/ritsource/episteme/prototype/data/models"
	"gopkg.in/yaml.v2"
)

func main() {
	bs, err := ioutil.ReadFile(constants.DEFAULT_DATA_INPUT_CONFIG_FILEPATH)
	if err != nil {
		fmt.Printf("unable to read file %v\n", constants.DEFAULT_DATA_INPUT_CONFIG_FILEPATH)
		panic(err)
	}

	err = os.MkdirAll(constants.DEFAULT_DATA_OUTPUT_DIR, 0700)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	type T struct {
		Posts          models.Posts          `json:"posts" yaml:"posts"`
		PinnedWebsites models.PinnedWebsites `json:"pinned_websites" yaml:"pinned_websites"`
	}

	data := T{}
	err = yaml.Unmarshal(bs, &data)
	// err = json.Unmarshal(bs, &data)
	if err != nil {
		fmt.Printf("unable to decode %v data\n", constants.DEFAULT_DATA_INPUT_CONFIG_FILEPATH)
		panic(err)
	}

	err = data.Posts.SaveToFS(constants.DEFAULT_POSTS_DATA_OUTPUT_FILEPATH)
	if err != nil {
		fmt.Printf("unable to save message data file to %+v\n", constants.DEFAULT_POSTS_DATA_OUTPUT_FILEPATH)
		panic(err)
	}

	err = data.PinnedWebsites.SaveToFS(constants.DEFAULT_PINNED_WEBSITES_DATA_OUTPUT_FILEPATH)
	if err != nil {
		fmt.Printf("unable to save message data file to %+v\n", constants.DEFAULT_PINNED_WEBSITES_DATA_OUTPUT_FILEPATH)
		panic(err)
	}
}
