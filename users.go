package retrieve_users

type Users []*User

func (users Users) SortedInsert(user *User) Users {
	length := len(users)
	i := 0

	// fast forward
	for i <= length-1 && users[i].Cmp(user) {
		i++
	}

	// insert the user
	if (i == 0) {
		// beginning
		users = append(Users{user}, users...)
	} else if (i == length) {
		// end
		users = append(users, user)
	} else {
		// middle
		first := users[:i]
		second := users[i:]

		users = append(first, append(Users{user}, second...)...)
	}

	return users
}