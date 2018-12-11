// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRunsParams creates a new GetRunsParams object
// with the default values initialized.
func NewGetRunsParams() *GetRunsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetRunsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunsParamsWithTimeout creates a new GetRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunsParamsWithTimeout(timeout time.Duration) *GetRunsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetRunsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetRunsParamsWithContext creates a new GetRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunsParamsWithContext(ctx context.Context) *GetRunsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetRunsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetRunsParamsWithHTTPClient creates a new GetRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunsParamsWithHTTPClient(client *http.Client) *GetRunsParams {
	var (
		pageNumberDefault = string("1")
		pageSizeDefault   = string("25")
	)
	return &GetRunsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetRunsParams contains all the parameters to send to the API endpoint
for the get runs operation typically these are written to a http.Request
*/
type GetRunsParams struct {

	/*ActivityID
	  The id of the activity.

	*/
	ActivityID *string
	/*EndTimeO2End
	  The upper bound of the timeO2End filter range.

	*/
	EndTimeO2End *strfmt.DateTime
	/*EndTimeO2Start
	  The upper bound of the timeO2Start filter range.

	*/
	EndTimeO2Start *strfmt.DateTime
	/*EndTimeTrgEnd
	  The upper bound of the timeTrgEnd filter range.

	*/
	EndTimeTrgEnd *strfmt.DateTime
	/*EndTimeTrgStart
	  The upper bound of the timeTrgStart filter range.

	*/
	EndTimeTrgStart *strfmt.DateTime
	/*OrderBy
	  On which field to order on.

	*/
	OrderBy *string
	/*OrderDirection
	  The order direction, either ascending or descending.

	*/
	OrderDirection *string
	/*PageNumber
	  The current page, i.e. the offset in the result set based on pageSize.

	*/
	PageNumber *string
	/*PageSize
	  The maximum amount of logs in each result.

	*/
	PageSize *string
	/*RunNumber
	  The id of the log.

	*/
	RunNumber *string
	/*RunQuality
	  The quality of the run.

	*/
	RunQuality *int64
	/*RunType
	  The type of the run.

	*/
	RunType *int64
	/*StartTimeO2End
	  The lower bound of the timeO2End filter range.

	*/
	StartTimeO2End *strfmt.DateTime
	/*StartTimeO2Start
	  The lower bound of the timeO2Start filter range.

	*/
	StartTimeO2Start *strfmt.DateTime
	/*StartTimeTrgEnd
	  The lower bound of the timeTrgEnd filter range.

	*/
	StartTimeTrgEnd *strfmt.DateTime
	/*StartTimeTrgStart
	  The lower bound of the timeTrgStart filter range.

	*/
	StartTimeTrgStart *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runs params
func (o *GetRunsParams) WithTimeout(timeout time.Duration) *GetRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runs params
func (o *GetRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runs params
func (o *GetRunsParams) WithContext(ctx context.Context) *GetRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runs params
func (o *GetRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) WithHTTPClient(client *http.Client) *GetRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActivityID adds the activityID to the get runs params
func (o *GetRunsParams) WithActivityID(activityID *string) *GetRunsParams {
	o.SetActivityID(activityID)
	return o
}

// SetActivityID adds the activityId to the get runs params
func (o *GetRunsParams) SetActivityID(activityID *string) {
	o.ActivityID = activityID
}

// WithEndTimeO2End adds the endTimeO2End to the get runs params
func (o *GetRunsParams) WithEndTimeO2End(endTimeO2End *strfmt.DateTime) *GetRunsParams {
	o.SetEndTimeO2End(endTimeO2End)
	return o
}

// SetEndTimeO2End adds the endTimeO2End to the get runs params
func (o *GetRunsParams) SetEndTimeO2End(endTimeO2End *strfmt.DateTime) {
	o.EndTimeO2End = endTimeO2End
}

// WithEndTimeO2Start adds the endTimeO2Start to the get runs params
func (o *GetRunsParams) WithEndTimeO2Start(endTimeO2Start *strfmt.DateTime) *GetRunsParams {
	o.SetEndTimeO2Start(endTimeO2Start)
	return o
}

// SetEndTimeO2Start adds the endTimeO2Start to the get runs params
func (o *GetRunsParams) SetEndTimeO2Start(endTimeO2Start *strfmt.DateTime) {
	o.EndTimeO2Start = endTimeO2Start
}

// WithEndTimeTrgEnd adds the endTimeTrgEnd to the get runs params
func (o *GetRunsParams) WithEndTimeTrgEnd(endTimeTrgEnd *strfmt.DateTime) *GetRunsParams {
	o.SetEndTimeTrgEnd(endTimeTrgEnd)
	return o
}

// SetEndTimeTrgEnd adds the endTimeTrgEnd to the get runs params
func (o *GetRunsParams) SetEndTimeTrgEnd(endTimeTrgEnd *strfmt.DateTime) {
	o.EndTimeTrgEnd = endTimeTrgEnd
}

// WithEndTimeTrgStart adds the endTimeTrgStart to the get runs params
func (o *GetRunsParams) WithEndTimeTrgStart(endTimeTrgStart *strfmt.DateTime) *GetRunsParams {
	o.SetEndTimeTrgStart(endTimeTrgStart)
	return o
}

// SetEndTimeTrgStart adds the endTimeTrgStart to the get runs params
func (o *GetRunsParams) SetEndTimeTrgStart(endTimeTrgStart *strfmt.DateTime) {
	o.EndTimeTrgStart = endTimeTrgStart
}

// WithOrderBy adds the orderBy to the get runs params
func (o *GetRunsParams) WithOrderBy(orderBy *string) *GetRunsParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get runs params
func (o *GetRunsParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the get runs params
func (o *GetRunsParams) WithOrderDirection(orderDirection *string) *GetRunsParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the get runs params
func (o *GetRunsParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithPageNumber adds the pageNumber to the get runs params
func (o *GetRunsParams) WithPageNumber(pageNumber *string) *GetRunsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get runs params
func (o *GetRunsParams) SetPageNumber(pageNumber *string) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get runs params
func (o *GetRunsParams) WithPageSize(pageSize *string) *GetRunsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get runs params
func (o *GetRunsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithRunNumber adds the runNumber to the get runs params
func (o *GetRunsParams) WithRunNumber(runNumber *string) *GetRunsParams {
	o.SetRunNumber(runNumber)
	return o
}

// SetRunNumber adds the runNumber to the get runs params
func (o *GetRunsParams) SetRunNumber(runNumber *string) {
	o.RunNumber = runNumber
}

// WithRunQuality adds the runQuality to the get runs params
func (o *GetRunsParams) WithRunQuality(runQuality *int64) *GetRunsParams {
	o.SetRunQuality(runQuality)
	return o
}

// SetRunQuality adds the runQuality to the get runs params
func (o *GetRunsParams) SetRunQuality(runQuality *int64) {
	o.RunQuality = runQuality
}

// WithRunType adds the runType to the get runs params
func (o *GetRunsParams) WithRunType(runType *int64) *GetRunsParams {
	o.SetRunType(runType)
	return o
}

// SetRunType adds the runType to the get runs params
func (o *GetRunsParams) SetRunType(runType *int64) {
	o.RunType = runType
}

// WithStartTimeO2End adds the startTimeO2End to the get runs params
func (o *GetRunsParams) WithStartTimeO2End(startTimeO2End *strfmt.DateTime) *GetRunsParams {
	o.SetStartTimeO2End(startTimeO2End)
	return o
}

// SetStartTimeO2End adds the startTimeO2End to the get runs params
func (o *GetRunsParams) SetStartTimeO2End(startTimeO2End *strfmt.DateTime) {
	o.StartTimeO2End = startTimeO2End
}

// WithStartTimeO2Start adds the startTimeO2Start to the get runs params
func (o *GetRunsParams) WithStartTimeO2Start(startTimeO2Start *strfmt.DateTime) *GetRunsParams {
	o.SetStartTimeO2Start(startTimeO2Start)
	return o
}

// SetStartTimeO2Start adds the startTimeO2Start to the get runs params
func (o *GetRunsParams) SetStartTimeO2Start(startTimeO2Start *strfmt.DateTime) {
	o.StartTimeO2Start = startTimeO2Start
}

// WithStartTimeTrgEnd adds the startTimeTrgEnd to the get runs params
func (o *GetRunsParams) WithStartTimeTrgEnd(startTimeTrgEnd *strfmt.DateTime) *GetRunsParams {
	o.SetStartTimeTrgEnd(startTimeTrgEnd)
	return o
}

// SetStartTimeTrgEnd adds the startTimeTrgEnd to the get runs params
func (o *GetRunsParams) SetStartTimeTrgEnd(startTimeTrgEnd *strfmt.DateTime) {
	o.StartTimeTrgEnd = startTimeTrgEnd
}

// WithStartTimeTrgStart adds the startTimeTrgStart to the get runs params
func (o *GetRunsParams) WithStartTimeTrgStart(startTimeTrgStart *strfmt.DateTime) *GetRunsParams {
	o.SetStartTimeTrgStart(startTimeTrgStart)
	return o
}

// SetStartTimeTrgStart adds the startTimeTrgStart to the get runs params
func (o *GetRunsParams) SetStartTimeTrgStart(startTimeTrgStart *strfmt.DateTime) {
	o.StartTimeTrgStart = startTimeTrgStart
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActivityID != nil {

		// query param activityId
		var qrActivityID string
		if o.ActivityID != nil {
			qrActivityID = *o.ActivityID
		}
		qActivityID := qrActivityID
		if qActivityID != "" {
			if err := r.SetQueryParam("activityId", qActivityID); err != nil {
				return err
			}
		}

	}

	if o.EndTimeO2End != nil {

		// query param endTimeO2End
		var qrEndTimeO2End strfmt.DateTime
		if o.EndTimeO2End != nil {
			qrEndTimeO2End = *o.EndTimeO2End
		}
		qEndTimeO2End := qrEndTimeO2End.String()
		if qEndTimeO2End != "" {
			if err := r.SetQueryParam("endTimeO2End", qEndTimeO2End); err != nil {
				return err
			}
		}

	}

	if o.EndTimeO2Start != nil {

		// query param endTimeO2Start
		var qrEndTimeO2Start strfmt.DateTime
		if o.EndTimeO2Start != nil {
			qrEndTimeO2Start = *o.EndTimeO2Start
		}
		qEndTimeO2Start := qrEndTimeO2Start.String()
		if qEndTimeO2Start != "" {
			if err := r.SetQueryParam("endTimeO2Start", qEndTimeO2Start); err != nil {
				return err
			}
		}

	}

	if o.EndTimeTrgEnd != nil {

		// query param endTimeTrgEnd
		var qrEndTimeTrgEnd strfmt.DateTime
		if o.EndTimeTrgEnd != nil {
			qrEndTimeTrgEnd = *o.EndTimeTrgEnd
		}
		qEndTimeTrgEnd := qrEndTimeTrgEnd.String()
		if qEndTimeTrgEnd != "" {
			if err := r.SetQueryParam("endTimeTrgEnd", qEndTimeTrgEnd); err != nil {
				return err
			}
		}

	}

	if o.EndTimeTrgStart != nil {

		// query param endTimeTrgStart
		var qrEndTimeTrgStart strfmt.DateTime
		if o.EndTimeTrgStart != nil {
			qrEndTimeTrgStart = *o.EndTimeTrgStart
		}
		qEndTimeTrgStart := qrEndTimeTrgStart.String()
		if qEndTimeTrgStart != "" {
			if err := r.SetQueryParam("endTimeTrgStart", qEndTimeTrgStart); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.OrderDirection != nil {

		// query param orderDirection
		var qrOrderDirection string
		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {
			if err := r.SetQueryParam("orderDirection", qOrderDirection); err != nil {
				return err
			}
		}

	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.RunNumber != nil {

		// query param runNumber
		var qrRunNumber string
		if o.RunNumber != nil {
			qrRunNumber = *o.RunNumber
		}
		qRunNumber := qrRunNumber
		if qRunNumber != "" {
			if err := r.SetQueryParam("runNumber", qRunNumber); err != nil {
				return err
			}
		}

	}

	if o.RunQuality != nil {

		// query param runQuality
		var qrRunQuality int64
		if o.RunQuality != nil {
			qrRunQuality = *o.RunQuality
		}
		qRunQuality := swag.FormatInt64(qrRunQuality)
		if qRunQuality != "" {
			if err := r.SetQueryParam("runQuality", qRunQuality); err != nil {
				return err
			}
		}

	}

	if o.RunType != nil {

		// query param runType
		var qrRunType int64
		if o.RunType != nil {
			qrRunType = *o.RunType
		}
		qRunType := swag.FormatInt64(qrRunType)
		if qRunType != "" {
			if err := r.SetQueryParam("runType", qRunType); err != nil {
				return err
			}
		}

	}

	if o.StartTimeO2End != nil {

		// query param startTimeO2End
		var qrStartTimeO2End strfmt.DateTime
		if o.StartTimeO2End != nil {
			qrStartTimeO2End = *o.StartTimeO2End
		}
		qStartTimeO2End := qrStartTimeO2End.String()
		if qStartTimeO2End != "" {
			if err := r.SetQueryParam("startTimeO2End", qStartTimeO2End); err != nil {
				return err
			}
		}

	}

	if o.StartTimeO2Start != nil {

		// query param startTimeO2Start
		var qrStartTimeO2Start strfmt.DateTime
		if o.StartTimeO2Start != nil {
			qrStartTimeO2Start = *o.StartTimeO2Start
		}
		qStartTimeO2Start := qrStartTimeO2Start.String()
		if qStartTimeO2Start != "" {
			if err := r.SetQueryParam("startTimeO2Start", qStartTimeO2Start); err != nil {
				return err
			}
		}

	}

	if o.StartTimeTrgEnd != nil {

		// query param startTimeTrgEnd
		var qrStartTimeTrgEnd strfmt.DateTime
		if o.StartTimeTrgEnd != nil {
			qrStartTimeTrgEnd = *o.StartTimeTrgEnd
		}
		qStartTimeTrgEnd := qrStartTimeTrgEnd.String()
		if qStartTimeTrgEnd != "" {
			if err := r.SetQueryParam("startTimeTrgEnd", qStartTimeTrgEnd); err != nil {
				return err
			}
		}

	}

	if o.StartTimeTrgStart != nil {

		// query param startTimeTrgStart
		var qrStartTimeTrgStart strfmt.DateTime
		if o.StartTimeTrgStart != nil {
			qrStartTimeTrgStart = *o.StartTimeTrgStart
		}
		qStartTimeTrgStart := qrStartTimeTrgStart.String()
		if qStartTimeTrgStart != "" {
			if err := r.SetQueryParam("startTimeTrgStart", qStartTimeTrgStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}