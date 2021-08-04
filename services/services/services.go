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

func (p *ServicesClient) Schdules(token, person_id string) (*SchedulesResponse, error) {
	httpReq, err := http.NewRequest(http.MethodGet, p.baseURL+fmt.Sprintf("/services/v2/people/%s/schedules", person_id), nil)
	if err != nil {
		return nil, err
	}

	resp := new(SchedulesResponse)

	return resp, utils.HttpRequestWithBearer(p.cl, httpReq, token, resp)
}
