package middleware
import (
	"github.com/gin-gonic/gin"
	"just.com/service/user"
	"encoding/json"
	"just.com/service/token"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := c.MustGet(MLEARNING_CONTENT).(*Context)
		apiKey := c.Query("api_key")
		userService := user.NewUserService(context.Session, context.Log)
		userTable, getFlag := userService.GetByEmail(apiKey)
		if getFlag == false {
			return
		}
		tokenService := service.NewTokenService(context.Session, context.Log)
		userToken, userTokenErr := tokenService.Make(userTable.UUID)
		if userTokenErr != nil {
			context.Log.Println(userTokenErr)
			return
		}
		tokenBytes, tokenErr := json.Marshal(userToken)
		if tokenErr != nil {
			context.Log.Println(tokenErr)
			return
		}
		c.Request.Header.Add(MLEARNING_HEADER_AUTHORIZATION, string(tokenBytes))
	}
}
