package retrieve_users

import (
	"testing"
)

func TestSortedInsertBeginning(t *testing.T) {
	// initial
	users := Users{&User{Name: "scorpion"}, &User{Name: "subzero"}}

	// insert another user (at the beginning)
	user := &User{Name: "liukang"}

	users = users.SortedInsert(user)

	// check it was inserted at the beginning
	if users[0] != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Name, users[0].Name)
	}
}

func TestSortedInsertEnd(t *testing.T) {
	// initial
	users := Users{&User{Name: "johnnycage"}, &User{Name: "raiden"}}

	// insert another user (at the end)
	user := &User{Name: "subzero"}

	users = users.SortedInsert(user)

	// check it was inserted at the end
	actual := users[len(users)-1]

	if actual != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Name, actual.Name)
	}
}

func TestSortedInsertMiddle(t *testing.T) {
	// initial
	users := Users{&User{Name: "johnnycage"}, &User{Name: "raiden"}}

	// insert another user (at the end)
	user := &User{Name: "kitana"}

	users = users.SortedInsert(user)

	// check it was inserted in the middle
	actual := users[1]

	if actual != user {
		t.Fatalf("Expected: %s. Actual: %s", user.Name, actual.Name)
	}
}
