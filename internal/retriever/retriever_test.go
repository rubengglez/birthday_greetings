package retriever_test

import (
	"testing"
	"testing/fstest"

	"github.com/rubengglez/birthday_greetings/internal/retriever"
	"github.com/rubengglez/birthday_greetings/internal/shared"
)

func TestFileSystemRetriever(t *testing.T) {
	assertFriend := func(t testing.TB, firstName, lastName, dateOfBirth, email string, friend shared.Friend) {
		t.Helper()
		if firstName != friend.FirstName {
			t.Errorf("got %s, want %s", friend.FirstName, firstName)
		}
		if lastName != friend.LastName {
			t.Errorf("got %s, want %s", friend.LastName, lastName)
		}
		if dateOfBirth != friend.DateOfBirth {
			t.Errorf("got %s, want %s", friend.DateOfBirth, dateOfBirth)
		}
		if email != friend.Email {
			t.Errorf("got %s, want %s", friend.Email, email)
		}
	}

	t.Run("Should get two friends from file", func(t *testing.T) {
		fs := fstest.MapFS{
			"friends.txt": {Data: []byte(`last_name, first_name, date_of_birth, email
Doe, John, 1982/10/08, john.doe@foobar.com
Ann, Mary, 1975/09/11, mary.ann@foobar.com`)},
		}

		retriever := retriever.NewFromFileSystem(fs)
		friends := retriever.Friends()

		if len(friends) != 2 {
			t.Errorf("got %d friends, want 2", len(friends))
		}
		assertFriend(t, "John", "Doe", "1982/10/08", "john.doe@foobar.com", friends[0])
		assertFriend(t, "Mary", "Ann", "1975/09/11", "mary.ann@foobar.com", friends[1])
	})

	t.Run("Should get no friends from file", func(t *testing.T) {
		fs := fstest.MapFS{
			"friends.txt": {Data: []byte("")},
		}

		retriever := retriever.NewFromFileSystem(fs)
		friends := retriever.Friends()

		if len(friends) != 0 {
			t.Errorf("got %d friends, want 0", len(friends))
		}
	})

	t.Run("Should get no friends when only first line is given from file", func(t *testing.T) {
		fs := fstest.MapFS{
			"friends.txt": {Data: []byte(`last_name, first_name, date_of_birth, email`)},
		}

		retriever := retriever.NewFromFileSystem(fs)
		friends := retriever.Friends()

		if len(friends) != 0 {
			t.Errorf("got %d friends, want 0", len(friends))
		}
	})
}
