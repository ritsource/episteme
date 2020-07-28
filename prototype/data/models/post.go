package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ritsource/episteme/prototype/data/protobuf/dst/postpb"
)

const POST_MSG_DST_DIR = ".data/posts"

// type Post struct {
// 	postpb.Post
// }

type Post = postpb.Post
type Post_Link = postpb.Post_Link
type Post_Link_Src = postpb.Post_Link_Src
type Post_Category = postpb.Post_Category

type Posts []*Post

func (ps Posts) SaveToFS(fp string) error {
	buf := new(bytes.Buffer)

	for _, p := range ps {
		if p.GetId() == "" {
			p.Id = uuid.New().String()
		}

		b, err := proto.Marshal(p)
		if err != nil {
			fmt.Printf("unable to read message (post.Id = %v)\n", p.GetId())
			return err
		}

		msglen := (len(b))

		buf.Grow(msglen + 4)

		err = binary.Write(buf, binary.LittleEndian, uint32(msglen))
		if err != nil {
			fmt.Printf("unable to read message (post.Id = %v)\n", p.GetId())
			return err
		}

		_, err = buf.Write(b)
		if err != nil {
			fmt.Printf("unable to read message (post.Id = %v)\n", p.GetId())
			return err
		}

	}

	if err := os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil && err != os.ErrExist {
		fmt.Printf("unable to create directory = %+v\n", filepath.Dir(fp))
		return err
	}

	file, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("unable to open file to write data")
		return err
	}
	defer file.Close()

	_, err = file.WriteAt(buf.Bytes(), 0)
	if err != nil {
		fmt.Println("unable to write data to file")
		return err
	}

	return nil
}

func (ps *Posts) ReadFromFS(fp string) (int, error) {
	n := 0
	posts := []*Post{}

	file, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("unable to open file %v\n", fp)
		return 0, err
	}
	defer file.Close()

	for {
		mlb := make([]byte, 4)

		n1, err := file.Read(mlb)
		n += n1
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("unable to read data from file\n")
			return n, err
		}

		msglen := binary.LittleEndian.Uint32(mlb)

		// buf := new(bytes.Buffer)
		// buf.Grow(msglen)
		buf := make([]byte, msglen)

		n2, err := file.Read(buf)
		n += n2
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("unable to read data from file\n")
			return n, err
		}

		post := Post{}
		err = proto.Unmarshal(buf, &post)
		if err != nil {
			fmt.Printf("failed to decode message from data\n")
			return n, err
		}

		posts = append(posts, &post)

	}

	// fmt.Printf("*** %+v\n", posts)
	// for _, post := range posts {
	// for	ps = append(ps, &post)
	// for}

	*ps = posts
	// fmt.Printf("*** %+v\n", ps)

	// isEq := reflect.DeepEqual(*ps, posts)
	// fmt.Printf("## ## ## ## isEq = %+v\n", isEq)

	return 0, nil
}
