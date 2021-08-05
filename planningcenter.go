package planningcenter

import (
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/oauth"
	"github.com/davidatdavidmarais/planningcenter/services/people"
	"github.com/davidatdavidmarais/planningcenter/services/services"
)

const (
	BaseURL = "https://api.planningcenteronline.com"
)

type PlanningCenter struct {
	OAuthClient    *oauth.OAuthClient
	PeopleClient   *people.PeopleClient
	ServicesClient *services.ServicesClient
}

func New(ClientID, ClientSecret string) *PlanningCenter {
	config := oauth.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		BaseURL:      BaseURL,
	}

	return &PlanningCenter{
		OAuthClient:    oauth.New(&http.Client{}, config),
		PeopleClient:   people.New(config.BaseURL, &http.Client{}),
		ServicesClient: services.New(config.BaseURL, &http.Client{}),
	}
}
