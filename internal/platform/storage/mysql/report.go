package mysql

import "time"

const (
	sqlReportsTable = "reports"
)

type sqlReport struct {
	StartDate        time.Time `db:"start_date"`
	EndDate          time.Time `db:"end_date"`
	JobAddress       string    `db:"job_address"`
	CodeCharge       string    `db:"code_charge"`
	Quantity         int32     `db:"quantity"`
	UnitsDescription string    `db:"units_description"`
}
