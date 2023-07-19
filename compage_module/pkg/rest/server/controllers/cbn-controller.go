package controllers

import (
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/models"
	"github.com/bheemeshkammak/compage_module/compage_module/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type CbnController struct {
	cbnService *services.CbnService
}

func NewCbnController() (*CbnController, error) {
	cbnService, err := services.NewCbnService()
	if err != nil {
		return nil, err
	}
	return &CbnController{
		cbnService: cbnService,
	}, nil
}

func (cbnController *CbnController) CreateCbn(context *gin.Context) {
	// validate input
	var input models.Cbn
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger cbn creation
	if _, err := cbnController.cbnService.CreateCbn(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Cbn created successfully"})
}

func (cbnController *CbnController) UpdateCbn(context *gin.Context) {
	// validate input
	var input models.Cbn
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbn update
	if _, err := cbnController.cbnService.UpdateCbn(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cbn updated successfully"})
}

func (cbnController *CbnController) FetchCbn(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbn fetching
	cbn, err := cbnController.cbnService.GetCbn(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cbn)
}

func (cbnController *CbnController) DeleteCbn(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger cbn deletion
	if err := cbnController.cbnService.DeleteCbn(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Cbn deleted successfully",
	})
}

func (cbnController *CbnController) ListCbns(context *gin.Context) {
	// trigger all cbns fetching
	cbns, err := cbnController.cbnService.ListCbns()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cbns)
}

func (*CbnController) PatchCbn(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*CbnController) OptionsCbn(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*CbnController) HeadCbn(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
