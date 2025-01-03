package handler

import (
	"devjudge/go-in-docker/model"
	"devjudge/go-in-docker/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var s = service.LeadService{}

func GetLeadById(ctx *gin.Context) {
	lead_id := ctx.Param("lead_id")
	leadId, err := strconv.Atoi(lead_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"reason": err.Error(),
		})
		return
	}
	resp := s.GetLeadByIdService(leadId)
	if resp == nil {
		ctx.JSON(http.StatusNotFound, gin.H{})
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}

func CreateLead(ctx *gin.Context) {

	var params model.Lead
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	lead, err := s.CreateLead(params)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"reason": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, lead)
	}
}
