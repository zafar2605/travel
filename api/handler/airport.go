package handler

import (
	"encoding/json"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateAirport godoc
// @ID create_airport
// @Router /airport [POST]
// @Summary Create Airport
// @Description Create Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.CreateAirport true "CreateAirportRequestBody"
// @Success 200 {object} Response{data=models.Airport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAirport(c *gin.Context) {
	var Airport = models.CreateAirport{}
	err := c.ShouldBindJSON(&Airport)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}

	resp, err := h.strg.Airport().Create(Airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Does not create"+err.Error())
		return
	}
	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdAirport godoc
// @ID get_by_id_airport
// @Router /airport/{id} [GET]
// @Summary Get By Id Airport
// @Description Get By Id Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Airport} "GetListAirportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportGetById(c *gin.Context) {
	var guid = c.Param("id")
	if !helpers.IsValidUUID(guid) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := h.strg.Airport().GetById(models.AirportPrimaryKey{Guid: guid})
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListAirport godoc
// @ID get_list_airport
// @Router /airport [GET]
// @Summary Get List Airport
// @Description Get List Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListAirportResponse} "GetListAirportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportGetList(c *gin.Context) {
	offset, err := h.getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, 400, "invalid offset")
		return
	}

	limit, err := h.getIntegerOrDefaultValue(c.Query("limit"), 0)
	if err != nil {
		handleResponse(c, 400, "invalid limit")
		return
	}

	resp, err := h.strg.Airport().GetList(models.GetListAirportRequest{
		Offset: int(offset),
		Limit:  int(limit),
	})
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateAirport godoc
// @ID update_airport
// @Router /airport [PUT]
// @Summary Update Airport
// @Description Update Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.UpdateAirport true "UpdateAirportRequestBody"
// @Success 200 {object} Response{data=models.Airport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportUpdate(c *gin.Context) {
	var Airport = models.UpdateAirport{}
	err := c.ShouldBindJSON(&Airport)

	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Update(Airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated")
}

// DeleteAirport godoc
// @ID delete_airport
// @Router /airport [DELETE]
// @Summary Delete Airport
// @Description Delete Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.AirportPrimaryKey true "DeleteAirportRequestBody"
// @Success 200 {object} Response{data=models.UpdateAirport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportDelete(c *gin.Context) {
	var Airport = models.AirportPrimaryKey{}
	err := c.ShouldBindJSON(&Airport)

	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Airport().Delete(Airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}

// UploadAirport godoc
// @ID upload_airport
// @Router /airport/:upload [POST]
// @Summary Upload airport
// @Description Upload Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param  	file  formData file true "File"
// @Success 200 {object} Response{data=[]models.CreateAirport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportUpload(c *gin.Context) {
	var airports = []models.CreateAirport{}
	file, err := c.FormFile("file")
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while file get "+err.Error())
		return
	}
	dts := file.Filename

	err = c.SaveUploadedFile(file, dts)
	defer os.Remove(dts)

	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, "Error while file save "+err.Error())
		return
	}

	body, _ := os.ReadFile(dts)
	err = json.Unmarshal(body, &airports)
	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, "Error read file "+err.Error())
		return
	}
	err = h.strg.Airport().Upload(airports)
	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, "Error while insert to postgres "+err.Error())
		return
	}
	handleResponse(c, http.StatusCreated, nil)
}
