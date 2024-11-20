package constant

import "net/http"

// Response message success
const (
	Success       = "Success"
	GetSuccess    = "Get Data Success"
	CreateSuccess = "Create Data Success"
	UpdateSuccess = "Update Data Success"
	DeleteSuccess = "Delete Data Success"
)

const (
	Error             = "Error"
	ErrDatabase       = "error database"
	ErrUnmarshal      = "Error unmarshal"
	ErrAlreadyExist   = "error already exist"
	ErrInvalidRequest = "invalid request"
	ErrNotFound       = "error not fount"
	ErrValidator      = "error validator"
	ErrGeneral        = "general error"
	ErrAuth           = "unathorized"
	ErrDefault        = "Something went wrong"
)

const (
	StatusOK                  = http.StatusOK
	StatusCreated             = http.StatusCreated
	StatusNotFound            = http.StatusNotFound
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusInternalServerError = http.StatusInternalServerError
)

const (
	DATA_TOKEN = "dataUserToken"
)
