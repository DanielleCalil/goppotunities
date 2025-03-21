package handler

import (
	"net/http"

	"github.com/DanielleCalil/goppotunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opennings")
		return
	}

	sendSuccess(ctx, "listOpenings", openings)
}
