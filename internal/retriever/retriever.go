package retriever

import "github.com/rubengglez/birthday_greetings/internal/shared"

type Retriever interface {
	Friends() shared.Friends
}

type MemoryRetriever struct{}

func (r *MemoryRetriever) Friends() shared.Friends {
	friends := make(shared.Friends, 0)

	friend := shared.Friend{
		LastName:    "Merlo",
		FirstName:   "Manolo",
		DateOfBirth: "1982/10/08",
		Email:       "example@dot.com",
	}
	friends = append(friends, friend)
	return friends
}
