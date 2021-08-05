package people

import (
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/utils"
)

const (
	PATH = "/people/v2/me"
)

type PeopleClient struct {
	baseURL string
	cl      *http.Client
}

func New(baseURL string, cl *http.Client) *PeopleClient {
	return &PeopleClient{
		baseURL: baseURL,
		cl:      cl,
	}
}

func (p *PeopleClient) Me(token string) (*Person, error) {
	httpReq, err := http.NewRequest(http.MethodGet, p.baseURL+PATH, nil)
	if err != nil {
		return nil, err
	}

	resp := new(Person)

	return resp, utils.HttpRequestWithBearer(p.cl, httpReq, token, resp)
}
