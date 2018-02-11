// Acceptance criteria
//
// We will wrap an imaginary database, represented by a slice, with our Proxy
// pattern. Then, the pattern will have to stick to the following acceptance
// criteria:
//
// 1. All accesse to the database of users will be done through the Proxy type.
//
// 2. A stack of n number of recent users will be kept in the Proxy.
//
// 3. If a user already exists in the stack, it won't query the database, and
// will return the stored one
//
// 4. If the queried user doesn't exist in the stack, it will query the
// database, remove the oldest user in the stack if it's full, store the new
// one, and return it.
//
package main

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		someDatabase = append(someDatabase, User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCapacity: 2,
		StackCache:    UserList{},
	}

	knowIDs := [3]int32{
		someDatabase[3].ID,
		someDatabase[4].ID,
		someDatabase[5].ID,
	}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knowIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knowIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search in an empty cache, the size of it " +
				"must be one")
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knowIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knowIDs[0] {
			t.Error("Returned user name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("Cache must not grouw if we ascked for an object that is stored on it")
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("The user should have been returned from cache")
		}
	})

	user1, err := proxy.FindUser(knowIDs[0])
	if err != nil {
		t.Fatal(err)
	}

	user2, _ := proxy.FindUser(knowIDs[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("The user wasn't stored on the proxy cache yet")
	}

	user3, _ := proxy.FindUser(knowIDs[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("The user wasn't stored on the proxy cache yet")
	}

	for i := 0; i < len(proxy.StackCache); i++ {
		if proxy.StackCache[i].ID == user1.ID {
			t.Error("User that should be gone was found")
		}
	}

	if len(proxy.StackCache) != 2 {
		t.Error("After inserting 3 users the cache should not grow more than two")
	}

	for _, v := range proxy.StackCache {
		if v != user2 && v != user3 {
			t.Error("A non expected user was found on the cache")
		}
	}
}
