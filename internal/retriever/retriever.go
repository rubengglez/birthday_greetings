package retriever

import "github.com/rubengglez/birthday_greetings/internal/shared"

type Retriever interface {
	Friends() shared.Friends
}

type MemoryRetriever struct{}

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
