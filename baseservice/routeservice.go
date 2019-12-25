package baseservice

/*
	ROUTE SERVICE
	@author Mufid Jamaluddin <mufidjamaluddin@outlook.com>
*/

import (
	"net/http"
	"encoding/json"
)

// RouteService ...
func RouteService(bs IBaseService) {
	var response []byte
	var err error

	var request *http.Request = bs.GetRequest()
	var responseWriter http.ResponseWriter = bs.GetWriter()

	responseWriter.Header().Set("Content-Type", "application/json")

	switch request.Method {
		case "GET":
			resultModel, resultError := bs.Get()
			if resultError == nil {
				responseWriter.WriteHeader(http.StatusOK)
				response, err = json.Marshal(resultModel)
			} else {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}
			break

		case "POST":
			resultModel, resultError := bs.Post()
			if resultError == nil {
				responseWriter.WriteHeader(http.StatusCreated)
				response, err = json.Marshal(resultModel)
			} else {
				responseWriter.WriteHeader(http.StatusConflict)
			}
			break

		case "PUT":
			resultModel, resultError := bs.Put()
			if resultError == nil {
        		responseWriter.WriteHeader(http.StatusAccepted)
				response, err = json.Marshal(resultModel)
			} else {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}
			break

		case "DELETE":
			resultModel, resultError := bs.Delete()
			if resultError == nil {
        		responseWriter.WriteHeader(http.StatusOK)
				response, err = json.Marshal(resultModel)
			} else {
				responseWriter.WriteHeader(http.StatusBadRequest)
			}
			break

    	default:
        	responseWriter.WriteHeader(http.StatusNotFound)
			response = []byte(`{"message": "not found"}`)
			break
	}
	
	if err == nil {
		responseWriter.Write(response)
	} else {
		panic(err)
	}
}