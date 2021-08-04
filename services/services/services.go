package services

import (
	"fmt"
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/utils"
)

type ServicesClient struct {
	baseURL string
	cl      *http.Client
}

func (s *ServicesClient) Schdules(token, person_id string) (*SchedulesResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, s.baseURL+fmt.Sprintf("/services/v2/people/%s/schedules", person_id), nil)
	if err != nil {
		return nil, err
	}

	resp := new(SchedulesResponse)

	return resp, utils.HttpRequestWithBearer(s.cl, httpReq, token, resp)
}

func (s *ServicesClient) Items(token, serviceTypeID, planID string) (*ItemsResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, s.baseURL+fmt.Sprintf("/services/v2/service_types/%s/plans/%s/items", serviceTypeID, planID), nil)
	if err != nil {
		return nil, err
	}

	resp := new(ItemsResponse)

	return resp, utils.HttpRequestWithBearer(s.cl, httpReq, token, resp)
}
