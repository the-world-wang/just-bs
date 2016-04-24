package user
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/middleware"
	"net/http"
	"just.com/service/user"
	"just.com/service/email"
	token_service "just.com/service/token"
)

type RegisterRequest struct {
	Email     string
	UserName  string `json:"user_name"`
	Password  string
	Password2 string
}

func RegisterHandle(c *gin.Context) {
	response := middleware.NewResponse(http.StatusOK, nil, nil)
	context, contextFlag := action.GetContext(c)
	if contextFlag == false {
		return
	}
	defer func() {
		context.Response = response
	}()

	request := new(RegisterRequest)
	bindErr := c.BindJSON(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	userService := user.NewUserService(context.Session, context.Log)
	userLoginVo, err := userService.Register(request.Email, request.UserName, request.Password, "STUDENT")
	if err != nil {
		context.Log.Println(err)
		return
	}
	response = middleware.NewResponse(http.StatusOK, userLoginVo, nil)
	go sendEmail(c, request.Email, request.UserName, userLoginVo.Token)
}

func sendEmail(c *gin.Context, email string, userName string, userToken *token_service.UserToken) {
	es := c.MustGet(middleware.MLEARNING__EMAIL).(*service.EmailService)
	es.SendMail(email, userName, userToken)
}