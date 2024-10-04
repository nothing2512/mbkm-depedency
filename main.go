package mbkmdepedency

type User struct {
	ID    int
	Name  string
	Email string
}

var users = []User{
	{
		ID:    1,
		Name:  "Fulan",
		Email: "fulan@gmail.com",
	},
	{
		ID:    2,
		Name:  "Fulanah",
		Email: "fulanah@gmail.com",
	},
	{
		ID:    3,
		Name:  "Ahmad",
		Email: "ahmad@gmail.com",
	},
}

// Create, Read, Update, Delete
func AddUser(user User) {
	users = append(users, user) //
}

func GetUsers() []User {
	return users
}

func GetUserByID(id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

func UpdateUser(id int, newuser User) {
	for index, user := range users {
		if user.ID == id {
			users[index] = newuser
		}
	}
}

func DeleteUser(id int) {
	for index, user := range users {
		if user.ID == id {
			// hapus user berdasarkan index
			users = append(users[:index], users[index+1:]...)
		}
	}
}
