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
	pinned_websitepb "github.com/ritsource/episteme/prototype/data/protobuf/dst/pinned_websitepb"
)

type PinnedWebsite = pinned_websitepb.PinnedWebsite
type PinnedWebsites []*PinnedWebsite

func (pws PinnedWebsites) SaveToFS(fp string) error {
	buf := new(bytes.Buffer)

	for _, w := range pws {
		if w.GetId() == "" {
			w.Id = uuid.New().String()
		}

		b, err := proto.Marshal(w)
		if err != nil {
			fmt.Printf("unable to read message\n")
			return err
		}

		msglen := (len(b))

		buf.Grow(msglen + 4)

		err = binary.Write(buf, binary.LittleEndian, uint32(msglen))
		if err != nil {
			fmt.Printf("unable to read message (Id = %v)\n", w.GetId())
			return err
		}

		_, err = buf.Write(b)
		if err != nil {
			fmt.Printf("unable to read message (Id = %v)\n", w.GetId())
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

func (pws *PinnedWebsites) ReadFromFS(fp string) (int, error) {
	n := 0
	websites := []*PinnedWebsite{}

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

		website := PinnedWebsite{}
		err = proto.Unmarshal(buf, &website)
		if err != nil {
			fmt.Printf("failed to decode message from data\n")
			return n, err
		}

		websites = append(websites, &website)

	}

	*pws = websites

	return 0, nil
}
