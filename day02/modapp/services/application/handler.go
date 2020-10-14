package shortcode

import (
	"encoding/json"
	"net/http"

	"github.com/danushkaf/go-basic-samples/day02/modapp/dto"
	"github.com/danushkaf/go-basic-samples/day02/modapp/util/common"
	"github.com/gorilla/mux"
)

func CreateApplication(s Service) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var applicationRequest dto.Application
		_ = json.NewDecoder(req.Body).Decode(&applicationRequest)

		code, message := common.ValidateApplicationCreateRequest(&applicationRequest)
		if code != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dto.Error{code, message})
			return
		}
		err := s.AddApplication(applicationRequest)
		common.SupportOptions(w, req)
		if (err != dto.ServiceError{}) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dto.Error{err.Code, "Request failed with message [" + err.Message + " ]"})
			return
		}
		json.NewEncoder(w).Encode(applicationRequest)
	}

}

func GetApplication(s Service) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		partnerId := params["partnerid"]
		applicationId := params["id"]

		partner, err := s.Get(partnerId, applicationId)
		common.SupportOptions(w, req)
		if (err != dto.ServiceError{}) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dto.Error{err.Code, "Request failed with message [" + err.Message + " ]"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(partner)
	}
}

func DeleteApplication(s Service) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		partnerId := params["partnerid"]
		applicationId := params["id"]

		err := s.Delete(partnerId, applicationId)
		common.SupportOptions(w, req)
		if (err != dto.ServiceError{}) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&dto.Error{err.Code, "Request failed with message [" + err.Message + " ]"})
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}

}
