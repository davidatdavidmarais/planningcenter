package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	httpReq.Header.Add("Authorization", "Bearer "+token)

	resp, err := p.cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	respObj := new(SchedulesResponse)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
