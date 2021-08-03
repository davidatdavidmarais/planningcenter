package oauth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/models"
)

const (
	EXCHANGE_TOKEN_GRANT_TYPE = "authorization_code"
	REFRESH_TOKEN_GRANT_TYPE  = "refresh_token"

	PATH = "/oauth/token"
)

type OAuthClient struct {
	config models.Config
	cl     *http.Client
}

func New(cl *http.Client, c models.Config) *OAuthClient {
	return &OAuthClient{
		cl:     cl,
		config: c,
	}
}

func (a *OAuthClient) ExchangeToken(code, redirectURI string) (*ExchangeTokenResponse, error) {
	req := &ExchangeTokenRequest{
		GrantType:    EXCHANGE_TOKEN_GRANT_TYPE,
		Code:         code,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		RedirectURI:  redirectURI,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(body)

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	cl := &http.Client{}
	resp, err := cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respObj := new(ExchangeTokenResponse)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

func (a *OAuthClient) RefreshToken(refreshToken string) (*ExchangeTokenResponse, error) {
	req := &RefreshTokenRequest{
		GrantType:    EXCHANGE_TOKEN_GRANT_TYPE,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		RefreshToken: refreshToken,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(body)

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	cl := &http.Client{}
	resp, err := cl.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respObj := new(ExchangeTokenResponse)
	err = json.Unmarshal(respBytes, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
