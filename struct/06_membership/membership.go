package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func newUser(name string, membershipType string) User {
	limit := 100
	if membershipType == "premium" {
		limit = 1000
	}

	return User{
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: limit,
		},
		Name: name,
	}
}
