package controllers

import (
	"net/http"
	//"encoding/json"
	"go-rest-api/dao"
	"github.com/julienschmidt/httprouter"
)

type PlanetsController struct{}

var dao = planetsDao{}

func NewPlanetsController() *PlanetsController {
	return &PlanetsController{}
}

func (dc PlanetsController) addPlanet(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	dao.Create()
	return nil, StatusOk(400)
}

func (dc PlanetsController) getAllPlanets(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	dao.GetAll()
	return nil, StatusOk(400)
}

func (dc PlanetsController) getPlanetById(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	dao.GetById()
	return nil, StatusOk(400)
}

func (dc PlanetsController) getPlanetByName(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	dao.GetByName()
	return nil, StatusOk(400)
}

func (dc PlanetsController) deletePlanetById(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) (interface{}, httpStatus) {
	dao.Delete()
	return nil, StatusOk(400)
}
