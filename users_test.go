package retrieve_users

import (
	"testing"
)

func TestSortedInsertBeginning(t *testing.T) {
	// initial
	users := Users{&User{Login: "scorpion"}, &User{Login: "subzero"}}

	// insert another user (at the beginning)
	user := &User{Login: "liukang"}

	users = users.SortedInsert(user)

	// check it was inserted at the beginning
	if users[0] != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Login, users[0].Login)
	}
}

func TestSortedInsertEnd(t *testing.T) {
	// initial
	users := Users{&User{Login: "johnnycage"}, &User{Login: "raiden"}}

	// insert another user (at the end)
	user := &User{Login: "subzero"}

	users = users.SortedInsert(user)

	// check it was inserted at the end
	actual := users[len(users)-1]

	if actual != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Login, actual.Login)
	}
}

func TestSortedInsertMiddle(t *testing.T) {
	// initial
	users := Users{&User{Login: "johnnycage"}, &User{Login: "raiden"}}

	// insert another user (at the end)
	user := &User{Login: "kitana"}

	users = users.SortedInsert(user)

	// check it was inserted in the middle
	actual := users[1]

	if actual != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Login, actual.Login)
	}
}
