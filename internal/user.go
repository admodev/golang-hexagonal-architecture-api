package report

import "context"

type User struct {
	id        int64
	token     string
	username  string
	email     string
	firstName string
	lastName  string
	website   string
	password  string
	role      string
}

func NewUser(token, username, email, firstName, lastName, website, password, role string) User {
	return User{
		token:     token,
		username:  username,
		email:     email,
		firstName: firstName,
		lastName:  lastName,
		website:   website,
		password:  password,
		role:      role,
	}
}

type UserRepository interface {
	Save(ctx context.Context, user User) error
}

func (u User) ID() int64 {
	return u.id
}

func (u User) Token() string {
	return u.token
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
	return u.website
}

func (u User) Password() string {
	return u.password
}

func (u User) Role() string {
	return u.role
}
