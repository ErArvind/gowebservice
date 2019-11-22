package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}
func AddUsers(u Users) (User, error) {
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return users, nil
}
