package report

import "context"

type User struct {
	username  string
	email     string
	firstName string
	lastName  string
	website   string
	password  string
	role      string
}

func NewUser(username, email, firstName, lastName, website, password, role string) User {
	if len(website) > 0 {
		return User{
			username:  username,
			email:     email,
			firstName: firstName,
			lastName:  lastName,
			website:   website,
			password:  password,
			role:      role,
		}
	} else {
		return User{
			username:  username,
			email:     email,
			firstName: firstName,
			lastName:  lastName,
			website:   "",
			password:  password,
			role:      role,
		}
	}
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

func (u User) Username() string {
	return u.username
}

func (u User) Email() string {
	return u.email
}

func (u User) FirstName() string {
	return u.firstName
}

func (u User) LastName() string {
	return u.lastName
}

func (u User) Website() string {
	if len(u.website) > 0 {
		return u.website
	} else {
		return "No website provided"
	}
}

func (u User) Password() string {
	return u.password
}

func (u User) Role() string {
	return u.role
}
