// Proxy
//
// It usually wraps an object to hide some of its characteristics. These
// characteristics could be the fact that it is a remote object (remote proxy),
// a very heavy object such as a very big image or the dump of a terabyte
// database (virtual proxy), or a restricted access object (protection proxy).
//
package main

import "fmt"

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found\n", id)
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

type UserListProxy struct {
	SomeDatabase           UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)

	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	fmt.Println("Returning user from database")
	u.DidLastSearchUsedCache = false
	u.addUserToStack(user)
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)

	} else {
		u.StackCache.addUser(user)
	}
}
