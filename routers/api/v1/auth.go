package v1

import (
	"kusnandartoni/starter/pkg/app"
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/services/svcmail"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterForm :
type RegisterForm struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email" valid:"Required"`
	Password        string `json:"password" valid:"Required"`
	FullName        string `json:"full_name,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MidleName       string `json:"midle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	PhotoURL        string `json:"photo_url,omitempty"`
	Gender          string `json:"gender,omitempty"`
	Birthday        string `json:"birthday,omitempty"`
	PhoneNumber     string `json:"phone_number,omitempty"`
	Address         string `json:"address,omitempty"`
	Industry        string `json:"industry,omitempty"`
	Company         string `json:"company,omitempty"`
	Occupation      string `json:"occupation,omitempty"`
	ExperienceLevel string `json:"experience_level,omitempty"`
	HighestDegree   string `json:"highest_degree,omitempty"`
	University      string `json:"university,omitempty"`
	Major           string `json:"major,omitempty"`
	Verified        bool   `json:"verified,omitempty"`
	UserRoles       string `json:"user_roles,omitempty"`
	Expertise       string `json:"expertise,omitempty"`
	Biography       string `json:"biography,omitempty"`
	Institution     string `json:"institution,omitempty"`
}

// Register :
// @Summary Register a Member
// @Tags Auth
// @Produce json
// @Param platform path string true "Platform"
// @Param req body v1.RegisterForm true "req param #changes are possible to adjust the form of the registration form from forntend"
// @Success 200 {object} app.Response
// @Router /api/auth/register [post]
func Register(c *gin.Context) {
	var (
		logger = logging.Logger{UUID: "0"}
		appG   = app.Gin{C: c}
		form   RegisterForm
		err    error
	)

	httpCode, errMsg := app.BindAndValid(c, &form)
	hashedPwd, err := util.Hash(form.Password)

	form.Password = hashedPwd
	logger.Info(form)

	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	httpCode, errMsg = registerMember(form)

	if httpCode > 0 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	mailService := svcmail.Verify{
		Email:      form.Email,
		UserName:   fmt.Sprintf("%s %s", form.FirstName, form.LastName),
		VerifyLink: fmt.Sprintf("%s/api/auth/verify?token=%s", setting.AppSetting.PrefixURL, util.GetEmailToken(form.Email)),
	}
	err = mailService.Store()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "OK", form))
}

// LoginForm :
type LoginForm struct {
	Email    string `valid:"Required"`
	Password string `valid:"Required"`
}

// Login :
// @Summary Login to get auth
// @Tags Auth
// @Produce  json
// @Param platform path string true "Platform"
// @Param req body v1.LoginForm true "req param"
// @Success 200 {object} app.Response
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var (
		logger = logging.Logger{UUID: "0"}
		appG   = app.Gin{C: c}
		form   LoginForm
		ID     int
	)

	httpCode, errMsg := app.BindAndValid(c, &form)
	// hashedPwd, _ := util.Hash(form.Password)

	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	httpCode, errMsg, ID = loginMember(form)

	if httpCode > 0 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	token, err := util.GenerateToken(ID)
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "OK", map[string]interface{}{
		"token": token,
		// "member": res,
	}))
}

// Verify :
// @Summary Verify email registration
// @Tags Auth
// @Produce  json
// @Param platform path string true "Platform"
// @Param token query string true "Token"
// @Success 200 {object} app.Response
// @Router /api/auth/verify [get]
func Verify(c *gin.Context) {
	var (
		logger   = logging.Logger{UUID: "0"}
		appG     = app.Gin{C: c}
		token    = c.Query("token")
		email    = util.ParseEmailToken(token)
		httpCode int
		errMsg   string
	)

	logger.Info(token)
	httpCode, errMsg = verifyMember(email)

	if httpCode > 0 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "OK", nil))
}

// ForgotForm :
type ForgotForm struct {
	Email string `valid:"Required"`
}

// Forgot :
// @Summary Forgot password
// @Tags Auth
// @Produce  json
// @Param platform path string true "Platform"
// @Param req body v1.ForgotForm true "req param"
// @Success 200 {object} app.Response
// @Router /api/auth/forgot [post]
func Forgot(c *gin.Context) {
	var (
		logger = logging.Logger{UUID: "0"}
		appG   = app.Gin{C: c}
		form   ForgotForm
		name   string
	)

	httpCode, errMsg := app.BindAndValid(c, &form)
	logger.Info(form)
	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	httpCode, errMsg, name = forgotMember(form.Email)
	if httpCode > 0 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	mailService := svcmail.ForgotPassword{
		UserName: name,
		Email:    form.Email,
		ResetLink: fmt.Sprintf("%s/api/auth/reset?token=%s",
			setting.AppSetting.PrefixURL,
			util.GetEmailToken(form.Email)),
	}

	err := mailService.Store()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "OK", form))
}

// ResetForm :
type ResetForm struct {
	Token    string `valid:"Required"`
	Password string `valid:"Required"`
}

// Reset :
// @Summary Reset email registration
// @Tags Auth
// @Produce  json
// @Param platform path string true "Platform"
// @Param req body v1.ResetForm true "req param"
// @Success 200 {object} app.Response
// @Router /api/auth/reset [put]
func Reset(c *gin.Context) {
	var (
		logger = logging.Logger{UUID: "0"}
		appG   = app.Gin{C: c}
		form   ResetForm
	)

	httpCode, errMsg := app.BindAndValid(c, &form)
	logger.Info(form)
	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	email := util.ParseEmailToken(form.Token)
	hashedPwd, _ := util.Hash(form.Password)

	httpCode, errMsg = resetMember(email, hashedPwd)

	if httpCode > 0 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "OK", nil))

}
