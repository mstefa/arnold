package external_sessions

import (
	"arnold/internal/external_login"
	gymdata "arnold/internal/gym"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExternalLoggingHandler(externalSessionService external_login.ExternalLooginService) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.Param("userid")
		err := externalSessionService.Loggin(ctx, userId)

		if err != nil {
			switch {
			case errors.Is(err, gymdata.ErrInvalidUserID):
				ctx.JSON(http.StatusBadRequest, err.Error()) //revisar otros errores
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
