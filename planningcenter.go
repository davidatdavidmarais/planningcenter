package planningcenter

import (
	"net/http"

	"github.com/davidatdavidmarais/planningcenter/services/models"
	"github.com/davidatdavidmarais/planningcenter/services/oauth"
)

const (
	BaseURL = "https://api.planningcenteronline.com"
)

type PlanningCenter struct {
	OAuthClient *oauth.OAuthClient
}

func New(ClientID, ClientSecret string) *PlanningCenter {
	config := models.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		BaseURL:      BaseURL,
	}

	return &PlanningCenter{
		OAuthClient: oauth.New(&http.Client{}, config),
	}
}
