package people

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	PATH = "/people/v2/me"
)

type PeopleClient struct {
	baseURL string
	cl      *http.Client
}

func (p *PeopleClient) Me(token string) (*Person, error) {
	httpReq, err := http.NewRequest(http.MethodGet, p.baseURL+PATH, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Authorization", "Bearer "+token)

	resp, err := p.cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("planning center api error")
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respObj := new(Person)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
