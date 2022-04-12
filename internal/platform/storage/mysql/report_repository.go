package mysql

import (
	report "bctec/internal"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

// ReportRepository is a MySQL report.ReportRepository implementation
type ReportRepository struct {
	db *sql.DB
}

// NewReportRepository inits a SQL-based implementation of report.ReportsRepository
func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{
		db: db,
	}
}

// Save implements the report.ReportRepository interface
func (r *ReportRepository) Save(ctx context.Context, report report.Report) error {
	reportSQLStruct := sqlbuilder.NewStruct(new(sqlReport))
	query, args := reportSQLStruct.InsertInto(sqlReportsTable, sqlReport{
		StartDate:        report.StartDate(),
		EndDate:          report.EndDate(),
		JobAddress:       report.JobAddress(),
		CodeCharge:       report.CodeCharge(),
		Quantity:         report.Quantity(),
		UnitsDescription: report.UnitsDescription(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error trying to persist report on database: %v", err)
	}

	return nil
}
