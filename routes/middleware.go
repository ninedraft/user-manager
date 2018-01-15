package routes

import (
	"net/http"

	"context"

	"git.containerum.net/ch/json-types/errors"
	umtypes "git.containerum.net/ch/json-types/user-manager"
	"git.containerum.net/ch/user-manager/server"
	"github.com/gin-gonic/gin"
)

type ReCaptchaRequest struct {
	ReCaptcha string `json:"recaptcha" binding:"required"`
}

var hdrToKey = map[string]interface{}{
	umtypes.UserIDHeader:      server.UserIDContextKey,
	"X-User-Agent":            server.UserAgentContextKey,
	umtypes.FingerprintHeader: server.FingerPrintContextKey,
	umtypes.SessionIDHeader:   server.SessionIDContextKey,
	"X-Token-ID":              server.TokenIDContextKey,
	"X-Client-IP":             server.ClientIPContextKey,
}

func prepareContext(ctx *gin.Context) {
	for hn, ck := range hdrToKey {
		if hv := ctx.GetHeader(hn); hv != "" {
			rctx := context.WithValue(ctx.Request.Context(), ck, hv)
			ctx.Request.WithContext(rctx)
		}
	}
}

func reCaptchaMiddleware(ctx *gin.Context) {
	/*var request ReCaptchaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	checkResp, err := svc.ReCaptchaClient.Check(ctx, ctx.ClientIP(), request.ReCaptcha)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, reCaptchaRequestFailed)
		return
	}

	if !checkResp.Success {
		ctx.AbortWithStatusJSON(http.StatusForbidden, errors.New(reCaptchaFailed))
	}*/
}

func errorWithHTTPStatus(err error) (int, *errors.Error) {
	switch err.(type) {
	case *server.AccessDeniedError:
		return http.StatusForbidden, err.(*server.AccessDeniedError).Err
	case *server.NotFoundError:
		return http.StatusNotFound, err.(*server.NotFoundError).Err
	case *server.BadRequestError:
		return http.StatusBadRequest, err.(*server.BadRequestError).Err
	case *server.AlreadyExistsError:
		return http.StatusConflict, err.(*server.AlreadyExistsError).Err
	case *server.InternalError:
		return http.StatusInternalServerError, err.(*server.InternalError).Err
	case *server.WebAPIError:
		return err.(*server.WebAPIError).StatusCode, err.(*server.WebAPIError).Err
	default:
		return http.StatusInternalServerError, errors.New(err.Error())
	}
}

func adminAccessMiddleware(ctx *gin.Context) {
	err := srv.CheckUserAdmin(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatusJSON(errorWithHTTPStatus(err))
	}
}
