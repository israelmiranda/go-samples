package users

import (
	"fmt"
	"sort"
)

type User struct {
	Username string
	Name     string
	Email    string
}

func NewUser(username, name, email string) *User {
	return &User{
		Username: username,
		Name:     name,
		Email:    email,
	}
}

func (u User) String() string {
	return fmt.Sprintf("User(username='%v', name='%v', email='%v')", u.Username, u.Name, u.Email)
}

func (u User) IntroduceYourself(guestName string) {
	fmt.Printf("Hi %v, I'm %v! Contact me at %v .\n", guestName, u.Name, u.Email)
}

type UserDatabase struct {
	Users []*User
}

func NewUserDatabase() *UserDatabase {
	return &UserDatabase{
		Users: []*User{},
	}
}

func (ud *UserDatabase) Insert(user *User) {
	i := sort.Search(len(ud.Users), func(i int) bool {
		return ud.Users[i].Username >= user.Username
	})
	ud.Users = append(ud.Users, nil)
	copy(ud.Users[i+1:], ud.Users[i:])
	ud.Users[i] = user
}

func (ud *UserDatabase) Find(username string) (*User, bool) {
	for _, user := range ud.Users {
		if user.Username == username {
			return user, true
		}
	}
	return nil, false
}

func (ud *UserDatabase) Update(user *User) {
	target, ok := ud.Find(user.Username)
	if ok {
		target.Name = user.Name
		target.Email = user.Email
	} else {
		fmt.Println("user not found")
	}
}

func (ud *UserDatabase) ListAll() []*User {
	return ud.Users
}

// func main() {
// user1 := new(User)
// fmt.Println(user1)

// user2 := NewUser("john", "John Doe", "john@doe.com")
// fmt.Println(user2)

// user3 := NewUser("jane", "Jane Doe", "jane@doe.com")
// user3.IntroduceYourself("David")

// user4 := NewUser("jane", "Jane Doe", "jane@doe.com")
// fmt.Println(user4)

// aakash := NewUser("aakash", "Aakash Rai", "aakash@example.com")
// biraj := NewUser("biraj", "Biraj Das", "biraj@example.com")
// hemanth := NewUser("hemanth", "Hemanth Jain", "hemanth@example.com")
// jadhesh := NewUser("jadhesh", "Jadhesh Verma", "jadhesh@example.com")
// siddhant := NewUser("siddhant", "Siddhant Sinha", "siddhant@example.com")
// sonaksh := NewUser("sonaksh", "Sonaksh Kumar", "sonaksh@example.com")
// vishal := NewUser("vishal", "Vishal Goel", "vishal@example.com")

// database := NewUserDatabase()

// database.Insert(biraj)
// database.Insert(hemanth)
// database.Insert(jadhesh)
// database.Insert(siddhant)
// database.Insert(sonaksh)
// database.Insert(vishal)
// database.Insert(aakash)

// fmt.Println(database.ListAll())
// }
