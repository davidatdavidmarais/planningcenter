package oauth

import (
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/utils"
)

const (
	EXCHANGE_TOKEN_GRANT_TYPE = "authorization_code"
	REFRESH_TOKEN_GRANT_TYPE  = "refresh_token"

	PATH = "/oauth/token"
)

type OAuthClient struct {
	config Config
	cl     *http.Client
}

func New(cl *http.Client, c Config) *OAuthClient {
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

	buf, err := utils.BodyBuffer(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	resp := new(ExchangeTokenResponse)

	return resp, utils.HttpRequest(a.cl, httpReq, resp)
}

func (a *OAuthClient) RefreshToken(refreshToken string) (*ExchangeTokenResponse, error) {
	req := &RefreshTokenRequest{
		GrantType:    REFRESH_TOKEN_GRANT_TYPE,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		RefreshToken: refreshToken,
	}

	buf, err := utils.BodyBuffer(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(http.MethodPost, a.config.BaseURL+PATH, buf)
	if err != nil {
		return nil, err
	}

	resp := new(ExchangeTokenResponse)

	return resp, utils.HttpRequest(a.cl, httpReq, resp)
}
