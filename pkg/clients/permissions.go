package clients

import (
	"context"
	"net/http"
	"time"

	"git.containerum.net/ch/auth/proto"
	"git.containerum.net/ch/user-manager/pkg/db"
	"github.com/containerum/cherry"
	headers "github.com/containerum/utils/httputil"
	utils "github.com/containerum/utils/httputil"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v1"
)

// PermissionsClient is an interface to resource-service.
type PermissionsClient interface {
	// GetUserAccess returns information about user access to resources (namespace, volumes) needed for token creation.
	GetUserAccess(ctx context.Context, user *db.User) (*authProto.ResourcesAccess, error)
	DeleteUserNamespaces(ctx context.Context, user *db.User) error
}

type httpPermissionsClient struct {
	rest *resty.Client
	log  *logrus.Entry
}

// NewHTTPPermissionsClient returns client for resource-service working via restful api
func NewHTTPPermissionsClient(serverURL string) PermissionsClient {
	log := logrus.WithField("component", "permissions_client")
	client := resty.New().
		SetHostURL(serverURL).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetTimeout(30 * time.Second).
		SetError(cherry.Err{})
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return &httpPermissionsClient{
		rest: client,
		log:  log,
	}
}

func (c *httpPermissionsClient) GetUserAccess(ctx context.Context, user *db.User) (*authProto.ResourcesAccess, error) {
	c.log.WithField("user_id", user.ID).Info("Getting user access from permissions service")
	headersMap := utils.RequestHeadersMap(ctx)
	headersMap[http.CanonicalHeaderKey(headers.UserIDXHeader)] = user.ID
	headersMap[http.CanonicalHeaderKey(headers.UserRoleXHeader)] = user.Role
	resp, err := c.rest.R().SetContext(ctx).
		SetResult(authProto.ResourcesAccess{}).
		SetHeaders(headersMap). // forward request headers to other our service
		Get("/accesses")
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, resp.Error().(*cherry.Err)
	}
	return resp.Result().(*authProto.ResourcesAccess), nil
}

func (c *httpPermissionsClient) DeleteUserNamespaces(ctx context.Context, user *db.User) error {
	c.log.WithField("user_id", user.ID).Info("Deleting user namespaces")
	headersMap := utils.RequestHeadersMap(ctx)
	headersMap[http.CanonicalHeaderKey(headers.UserIDXHeader)] = user.ID
	headersMap[http.CanonicalHeaderKey(headers.UserRoleXHeader)] = user.Role
	resp, err := c.rest.R().SetContext(ctx).
		SetResult(authProto.ResourcesAccess{}).
		SetHeaders(headersMap). // forward request headers to other our service
		Delete("/namespaces")
	if err != nil {
		return err
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}
