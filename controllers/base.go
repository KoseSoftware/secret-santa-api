package controllers

type BaseController struct {
	dbConn string
}

func NewBaseController(dbConn string) *BaseController {
	return &BaseController{
		dbConn: dbConn,
	}
}
