package clients

import (
	"net/http"

	umtypes "git.containerum.net/ch/json-types/user-manager"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

type WebAPIClient struct {
	log    *logrus.Entry
	client *resty.Client
}

func NewWebAPIClient(serverUrl string) *WebAPIClient {
	log := logrus.WithField("component", "web_api_client")
	client := resty.New().
		SetHostURL(serverUrl).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetHeader("Content-Type", "application/json").
		SetError(umtypes.WebAPIError{})
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return &WebAPIClient{
		log:    log,
		client: client,
	}
}

// returns raw answer from web-api
func (c *WebAPIClient) Login(request *umtypes.WebAPILoginRequest) (ret map[string]interface{}, statusCode int, err error) {
	c.log.WithField("login", request.Username).Infoln("Signing in through web-api")

	ret = make(map[string]interface{})
	resp, err := c.client.R().SetBody(request).SetResult(&ret).Post("/api/login")
	if err != nil {
		c.log.WithError(err).Errorln("Sign in through web-api request failed")
		return nil, http.StatusInternalServerError, err
	}
	if resp.Error() != nil {
		err = resp.Error().(*umtypes.WebAPIError)
	}

	return ret, resp.StatusCode(), err
}