package reports

import (
	reports "bctec/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createRequest struct {
	StartDate        time.Time `json:"start_date" binding:"required,field=EndDate" time_format:"2022-11-04"`
	EndDate          time.Time `json:"end_date" binding:"required" time_format:"2022-11-04"`
	JobAddress       string    `json:"job_address" binding:"required"`
	CodeCharge       string    `json:"code_charge" binding:"required"`
	Quantity         int32     `json:"quantity" binding:"required"`
	UnitsDescription string    `json:"units_description" binding:"required"`
}

// CreateHandler returns an HTTP handler for reports creation
func CreateHandler(reportsRepository reports.ReportsRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		report := reports.NewReport(req.StartDate, req.EndDate, req.JobAddress, req.CodeCharge, req.Quantity, req.UnitsDescription)

		if err := reportsRepository.Save(ctx, report); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
