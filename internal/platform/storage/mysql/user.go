package mysql

const (
	sqlUsersTable = "users"
)

type sqlUser struct {
	Token     string `db:"token"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Website   string `db:"website"`
	Password  string `db:"password"`
	Role      string `db:"role"`
}
