package report

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
