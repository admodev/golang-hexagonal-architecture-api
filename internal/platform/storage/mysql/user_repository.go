package mysql

import (
	"bctec/internal"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

// UserRepository is a MySQL report.ReportRepository implementation
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository inits a SQL-based implementation of report.ReportsRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save implements the users.ReportRepository interface
func (r *UserRepository) Save(ctx context.Context, user report.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(sqlUser))
	query, args := userSQLStruct.InsertInto(sqlUsersTable, sqlUser{
		Token:     user.Token(),
		Username:  user.Username(),
		Email:     user.Email(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Website:   user.Website(),
		Password:  user.Password(),
		Role:      user.Role(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error trying to persist report on database: %v", err)
	}

	return nil
}
