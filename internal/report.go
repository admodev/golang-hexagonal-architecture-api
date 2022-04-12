package report

import (
	"context"
	"time"
)

// Report is the data structure that represents a report
type Report struct {
	startDate        time.Time
	endDate          time.Time
	jobAddress       string
	codeCharge       string
	quantity         int32
	unitsDescription string
}

func NewReport(startDate time.Time, endDate time.Time, jobAddress string, codeCharge string, quantity int32, unitsDescription string) Report {
	return Report{
		startDate:        startDate,
		endDate:          endDate,
		jobAddress:       jobAddress,
		codeCharge:       codeCharge,
		quantity:         quantity,
		unitsDescription: unitsDescription,
	}
}

// ReportsRepository defines the expected behaviour from a report storage
type ReportsRepository interface {
	Save(ctx context.Context, report Report) error
}

func (r Report) StartDate() time.Time {
	return r.startDate
}

func (r Report) EndDate() time.Time {
	return r.endDate
}

func (r Report) JobAddress() string {
	return r.jobAddress
}

func (r Report) CodeCharge() string {
	return r.codeCharge
}

func (r Report) Quantity() int32 {
	return r.quantity
}

func (r Report) UnitsDescription() string {
	return r.unitsDescription
}
