package common

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	dto "github.com/danushkaf/go-basic-samples/day02/modapp/dto"
	"github.com/danushkaf/go-basic-samples/day02/modapp/util"
	"github.com/google/uuid"
)

func GetPaginationVars(req *http.Request) dto.PaginationStruct {
	vals := req.URL.Query()
	var offset string
	var limit int
	var isPaginationRequested bool
	var error error
	if len(vals) > 0 {
		offsets, haveOffset := vals["offset"]
		if haveOffset {
			if len(offsets) >= 1 {
				offset = offsets[0]
				isPaginationRequested = true
			}
		} else {
			isPaginationRequested = false
		}
		limits, haveLimit := vals["limit"]
		if haveLimit {
			if len(limits) >= 1 {
				limit, error = strconv.Atoi(limits[0])
				if error == nil {
					isPaginationRequested = true
				} else {
					isPaginationRequested = false
				}
			}
		} else {
			isPaginationRequested = false
		}
	} else {
		isPaginationRequested = false
	}
	return dto.PaginationStruct{
		Offset:                offset,
		Limit:                 limit,
		IsPaginationRequested: isPaginationRequested,
	}
}

func ConvertStringToInt(stringValue string) int {
	val, err := strconv.Atoi(stringValue)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return val
}

func ConvertIntToString(intValue int) string {
	return strconv.Itoa(intValue)
}

func SupportOptions(w http.ResponseWriter, req *http.Request) {
	SetupResponse(&w, req)
}

func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func GetRandonUUID() string {
	return uuid.New().String()
}

func GetProperty(propertyName string) string {
	return util.GetProperties().MustGetString(propertyName)
}

func GetDBUrls() []string {
	var nodes []string
	nodeCount := ConvertStringToInt(GetProperty("cassandraNodeCount"))
	for i := 0; i < nodeCount; i++ {
		nodeId := i + 1
		nodes = append(nodes, GetProperty("cassandra-node-"+ConvertIntToString(nodeId)))
	}
	return nodes
}

func ValidateApplicationCreateRequest(application *dto.Application) (string, string) {

	minApplicationNameLength := ConvertStringToInt(GetProperty("applicationName-length-min"))
	maxApplicationNameLength := ConvertStringToInt(GetProperty("applicationName-length-max"))
	if application.Name == "" {
		return "10001", "Missing application name"
	}

	if application.PartnerID == "" {
		return "10001", "Missing application partner id"
	}

	if len(application.Name) < minApplicationNameLength || len(application.Name) > maxApplicationNameLength {
		return "10002", "Parameter not within expected length"
	}

	if strings.ContainsAny(application.Name, GetProperty("special-characters")) {
		return "10005", "Contains special characters"
	}

	var applicationId = application.ID
	if applicationId != "" {
		if strings.ContainsAny(applicationId, GetProperty("special-characters")) {
			return "10005", "Contains special characters"
		}

		minApplicationIdLength := ConvertStringToInt(GetProperty("applicationId-length-min"))
		maxApplicationIdLength := ConvertStringToInt(GetProperty("applicationId-length-max"))

		if len(applicationId) < minApplicationIdLength || len(applicationId) > maxApplicationIdLength {
			return "10002", "Parameter not within expected length"
		}
	}

	if application.ID == "" {
		application.ID = GetRandonUUID()
	}

	strings.TrimSpace(application.Name)
	strings.TrimSpace(application.ID)

	return "", ""
}
