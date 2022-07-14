package dto

import "net/http"

// Statuses
const (
	STATUS_FAIL    = 666
	STATUS_OK      = http.StatusOK
	STATUS_BAD_REQ = http.StatusBadRequest
)

// Parameters
const (
	PARAM_INT_MIN = "min"
	PARAM_INT_MAX = "max"
	PARAM_PRAYER  = "num"
)
