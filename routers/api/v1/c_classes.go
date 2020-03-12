package v1

import (
	"fmt"
	"kusnandartoni/starter/pkg/app"
	"kusnandartoni/starter/pkg/logging"
	"kusnandartoni/starter/pkg/setting"
	"kusnandartoni/starter/pkg/util"
	"kusnandartoni/starter/services/svcclasses"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// GetClasses :
// @Summary Get Classes
// @Security ApiKeyAuth
// @Tags MST Classes
// @Produce  json
// @Param id query int false "ID"
// @Param page query int false "Page"
// @Success 200 {object} app.Response
// @Router /api/v1/class [get]
func GetClasses(c *gin.Context) {
	var (
		logger = logging.Logger{}
		err    = mapstructure.Decode(app.GetClaims(c), &logger)
		appG   = app.Gin{C: c}
		id     = int64(-1)
	)

	if arg := c.Query("id"); arg != "" {
		id = int64(com.StrTo(arg).MustInt())
	}

	classesService := svcclasses.Classes{
		ID:       id,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	logger.Info(util.Stringify(classesService))
	classes, err := classesService.GetAll()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	count, err := classesService.Count()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "Ok", map[string]interface{}{
		"list":  classes,
		"total": count,
	}))
}

// AddClassForm :
type AddClassForm struct {
	ImageURL    string `json:"image_url" valid:"Required"`
	Name        string `json:"name" valid:"Required"`
	Description string `json:"description,omitempty"`
	Headline    string `json:"headline,omitempty"`
}

// AddClass :
// @Summary AddClass
// @Security ApiKeyAuth
// @Tags MST Classes
// @Produce  json
// @Param req body v1.AddClassForm true "req param"
// @Success 200 {object} app.Response
// @Router /api/v1/class [post]
func AddClass(c *gin.Context) {
	var (
		logger         = logging.Logger{}
		err            = mapstructure.Decode(app.GetClaims(c), &logger)
		appG           = app.Gin{C: c}
		classesService svcclasses.Classes
		form           AddClassForm
	)

	httpCode, errMsg := app.BindAndValid(c, &form)
	logger.Info(util.Stringify(form))

	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	err = mapstructure.Decode(form, &classesService)
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}
	classesService.CreatedBy = logger.UUID
	classesService.ModifiedBy = logger.UUID
	err = classesService.Add()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "Data berhasil ditambah", classesService))
}

// EditClassForm :
type EditClassForm struct {
	ImageURL    string `json:"image_url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Headline    string `json:"headline,omitempty"`
}

// EditClass :
// @Summary Edit Class
// @Security ApiKeyAuth
// @Tags MST Classes
// @Produce  json
// @Param id path int true "ID"
// @Param req body v1.EditClassForm true "req param"
// @Success 200 {object} app.Response
// @Router /api/v1/class/{id} [put]
func EditClass(c *gin.Context) {
	var (
		logger         = logging.Logger{}
		err            = mapstructure.Decode(app.GetClaims(c), &logger)
		appG           = app.Gin{C: c}
		id             = int64(com.StrTo(c.Param("id")).MustInt())
		form           EditClassForm
		classesService svcclasses.Classes
		valid          validation.Validation
	)

	valid.Min(id, 1, "id").Message("ID must be greater than 0")
	if valid.HasErrors() {
		logger.Error(appG.Response(http.StatusBadRequest, app.MarkErrors(valid.Errors), nil))
		return
	}

	httpCode, errMsg := app.BindAndValid(c, &form)
	logger.Info(util.Stringify(form))
	if httpCode != 200 {
		logger.Error(appG.Response(httpCode, errMsg, nil))
		return
	}

	err = mapstructure.Decode(form, &classesService)
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	classesService.ID = id
	_, err = classesService.ExistByID()
	if err != nil {
		logger.Error(appG.Response(http.StatusUnprocessableEntity, fmt.Sprintf("%v", err), nil))
		return
	}

	classesService.ModifiedBy = logger.UUID
	err = classesService.Edit()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}
	logger.Info(appG.Response(http.StatusOK, "Data berhasil diubah", form))

}

// DeleteClass :
// @Summary Delete  Class
// @Security ApiKeyAuth
// @Tags MST Classes
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Router /api/v1/class/{id} [delete]
func DeleteClass(c *gin.Context) {
	var (
		logger         = logging.Logger{}
		err            = mapstructure.Decode(app.GetClaims(c), &logger)
		appG           = app.Gin{C: c}
		id             = int64(com.StrTo(c.Param("id")).MustInt())
		valid          validation.Validation
		classesService svcclasses.Classes
	)

	valid.Min(id, 1, "id").Message("ID must be greater than 0")
	logger.Info(id)
	if valid.HasErrors() {
		logger.Error(appG.Response(http.StatusBadRequest, app.MarkErrors(valid.Errors), nil))
		return
	}

	classesService.ID = id
	_, err = classesService.ExistByID()
	if err != nil {
		logger.Error(appG.Response(http.StatusUnprocessableEntity, fmt.Sprintf("%v", err), nil))
		return
	}

	classesService.DeletedBy = logger.UUID
	err = classesService.Delete()
	if err != nil {
		logger.Error(appG.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil))
		return
	}

	logger.Info(appG.Response(http.StatusOK, "Data berhasil dihapus", id))
}
