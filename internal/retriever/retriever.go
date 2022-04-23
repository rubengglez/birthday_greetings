package retriever

import (
	"bufio"
	"io/fs"
	"log"
	"strings"

	"github.com/rubengglez/birthday_greetings/internal/shared"
)

type Retriever interface {
	Friends() shared.Friends
}

type MemoryRetriever struct{}
type FileRetriever struct {
	file string
	fs   fs.FS
}

func (r *MemoryRetriever) Friends() shared.Friends {
	friends := make(shared.Friends, 0)

	manolo := shared.Friend{
		LastName:    "Merlo",
		FirstName:   "Manolo",
		DateOfBirth: "1982/10/08",
		Email:       "manolo@dot.com",
	}
	paco := shared.Friend{
		LastName:    "Merlo",
		FirstName:   "Paco",
		DateOfBirth: "1982/04/03",
		Email:       "paco@dot.com",
	}
	friends = append(friends, manolo)
	friends = append(friends, paco)
	return friends
}

func NewFromFileSystem(fileSystem fs.FS) Retriever {
	return &FileRetriever{file: "friends.txt", fs: fileSystem}
}

func (r *FileRetriever) Friends() shared.Friends {
	log.Println(r.fs)
	fileHandle, err := r.fs.Open("friends.txt")
	if err != nil {
		panic("oppss")
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var friends = make(shared.Friends, 0)

	i := 0
	for fileScanner.Scan() {
		i++
		if i == 1 {
			continue
		}
		data := strings.Split(fileScanner.Text(), ",")

		friend := shared.Friend{
			LastName:    strings.TrimSpace(data[0]),
			FirstName:   strings.TrimSpace(data[1]),
			DateOfBirth: strings.TrimSpace(data[2]),
			Email:       strings.TrimSpace(data[3]),
		}

		friends = append(friends, friend)
	}

	return friends
}
