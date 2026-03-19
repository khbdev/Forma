package handler

import (
	"net/http"
	"strconv"

	"forma/internal/domain"
	"forma/internal/models"
	"forma/pkg/response"

	"github.com/gin-gonic/gin"
)

type LeadHandler struct {
	service domain.LeadService
}

func NewLeadHandler(service domain.LeadService) *LeadHandler {
	return &LeadHandler{
		service: service,
	}
}

func (h *LeadHandler) Create(c *gin.Context) {
	var lead models.Lead

	if err := c.ShouldBindJSON(&lead); err != nil {
		response.Error(c.Writer, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if err := h.service.Create(c.Request.Context(), &lead); err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	response.Success(
		c.Writer,
		http.StatusCreated,
		"Arizangiz qabul qilindi. Siz bilan yaqin orada bog‘lanamiz.",
	)
}

func (h *LeadHandler) GetAll(c *gin.Context) {
	leads, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	// bu yerda data qaytaramiz (response package faqat message uchun)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    leads,
	})
}

func (h *LeadHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, "Noto‘g‘ri ID", nil)
		return
	}

	lead, err := h.service.GetByID(c.Request.Context(), uint(id64))
	if err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	if lead == nil {
		response.Error(c.Writer, http.StatusNotFound, "Lead topilmadi", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lead,
	})
}

func (h *LeadHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		response.Error(c.Writer, http.StatusBadRequest, "Noto‘g‘ri ID", nil)
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id64)); err != nil {
		response.Error(c.Writer, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	response.Success(c.Writer, http.StatusOK, "Lead muvaffaqiyatli o‘chirildi")
}