package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func Include() *httprouter.Router {

	router := httprouter.New()

	planetsController := NewPlanetsController()

	router.POST("/planet", PublicRoute(planetsController.addPlanet))
	router.GET("/planet/all", PublicRoute(planetsController.getAllPlanets))
	router.GET("/planet/id/:id", PublicRoute(planetsController.getPlanetById))
	router.GET("/planet/name/:name", PublicRoute(planetsController.getPlanetByName))
	router.DELETE("/planet/:id", PublicRoute(planetsController.deletePlanetById))

	return router
}

type controllerRoute func(http.ResponseWriter, *http.Request, httprouter.Params) (interface{}, httpStatus)

func PublicRoute(controller controllerRoute) httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {

		result, status := controller(responseWriter, request, params)
		if status.err != nil {
			WriteResponse(responseWriter, nil, ServerError(nil))
		} else {
			WriteResponse(responseWriter, result, status)
		}

	}
}

func WriteResponse(responseWriter http.ResponseWriter, result interface{}, httpStatus httpStatus) {

	var responseBody string
	var responseApi ResponseApi

	responseWriter.Header().Set("Content-Type", "application/json")
	if httpStatus.err != nil {
		responseWriter.WriteHeader(httpStatus.status)
		responseApi.Type = "ERROR"
		responseApi.Data = httpStatus.err.Error()
		jsonBody, _ := json.Marshal(responseApi)
		responseBody = string(jsonBody)
	} else {
		responseWriter.WriteHeader(httpStatus.status)
		responseApi.Type = "SUCCESS"
		responseApi.Data = result
		jsonBody, _ := json.Marshal(responseApi)
		responseBody = string(jsonBody)
	}
	fmt.Fprintf(responseWriter, responseBody)
}

type ResponseApi struct {
	Type     string `json:"type"`
	Data     interface{} `json:"data"`
  Messages []string `json:"messages"`
}

type httpStatus struct {
	err    error
	status int
}

func ServerError(err error) httpStatus {
	return httpStatus{err, http.StatusInternalServerError}
}

func StatusOk(status int) httpStatus {
	return httpStatus{nil, status}
}

func StatusNotFound(err error) httpStatus {
	return httpStatus{err, http.StatusBadRequest}
}

