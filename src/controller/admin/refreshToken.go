package admin_controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	admin_controller_util "github.com/matheuswww/mystream/src/controller/admin/util"
	"github.com/matheuswww/mystream/src/logger"
	rest_err "github.com/matheuswww/mystream/src/restErr"
)

func (ac *adminController) RefreshToken(c *gin.Context) {
	logger.Log("Init RefreshToken")
	authHeader := c.GetHeader("Authorization")
	refreshToken, err := admin_controller_util.GetToken(authHeader)
	if err != nil {
		restErr := rest_err.NewBadRequestError(err.Error())
		c.JSON(restErr.Code, restErr)
		return
	}
	token,restErr := ac.admin_service.RefreshToken(refreshToken)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, token)
}