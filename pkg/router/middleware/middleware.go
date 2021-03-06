package middleware

import (
	"net/textproto"

	"git.containerum.net/ch/user-manager/pkg/server"
	"git.containerum.net/ch/user-manager/pkg/umerrors"
	"github.com/containerum/cherry/adaptors/gonic"
	headers "github.com/containerum/utils/httputil"
	"github.com/gin-gonic/gin"
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

// RequireAdminRole
func RequireAdminRole(ctx *gin.Context) {
	if ctx.GetHeader(textproto.CanonicalMIMEHeaderKey(headers.UserRoleXHeader)) != RoleAdmin {
		gonic.Gonic(umerrors.ErrAdminRequired(), ctx)
		return
	}

	um := ctx.MustGet(UMServices).(server.UserManager)
	err := um.CheckAdmin(ctx.Request.Context())
	if err != nil {
		gonic.Gonic(umerrors.ErrAdminRequired(), ctx)
	}
}

func RequireUserExist(ctx *gin.Context) {
	um := ctx.MustGet(UMServices).(server.UserManager)
	err := um.CheckUserExist(ctx.Request.Context())
	if err != nil {
		gonic.Gonic(umerrors.ErrUserNotExist(), ctx)
	}
}
