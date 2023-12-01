package handler

import (
	"encoding/json"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// CreateCountry godoc
// @ID create_country
// @Router /country [POST]
// @Summary Create Country
// @Description Create Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CreateCountry true "CreateCountryRequestBody"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCountry(c *gin.Context) {
	var country = models.CreateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}

	resp, err := h.strg.Country().Create(country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Does not create"+err.Error())
		return
	}
	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCountrygodoc
// @ID get_by_id_country
// @Router /country/{id} [GET]
// @Summary Get By Id Country
// @Description Get By Id Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Country} "GetListCountryResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryGetById(c *gin.Context) {
	var guid = c.Param("guid")
	if !helpers.IsValidUUID(guid) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	resp, err := h.strg.Country().GetById(models.CountryPrimaryKey{Guid: guid})
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListCountrygodoc
// @ID get_list_country
// @Router /country [GET]
// @Summary Get List Country
// @Description Get List Country
// @Tags Country
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListCountryResponse} "GetListCountryResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryGetList(c *gin.Context) {
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

	resp, err := h.strg.Country().GetList(models.GetListCountryRequest{
		Offset: int(offset),
		Limit:  int(limit),
	})
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateCountry godoc
// @ID update_country
// @Router /country [PUT]
// @Summary Update Country
// @Description Update Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.UpdateCountry true "UpdateCountryRequestBody"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryUpdate(c *gin.Context) {
	var country = models.UpdateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Country().Update(country)
	if err != nil {
		handleResponse(c, 500, "Country does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Updated")
}

// DeleteCountry godoc
// @ID delete_country
// @Router /country [DELETE]
// @Summary Delete Country
// @Description Delete Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CountryPrimaryKey true "DeleteCountryRequestBody"
// @Success 200 {object} Response{data=models.UpdateCountry} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryDelete(c *gin.Context) {
	var country = models.CountryPrimaryKey{}
	err := c.ShouldBindJSON(&country)

	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	_, err = h.strg.Country().Delete(country)
	if err != nil {
		handleResponse(c, 500, "Country does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}

// UploadCountry godoc
// @ID upload_country
// @Router /country/:upload [POST]
// @Summary Upload country
// @Description Upload Country
// @Tags Country
// @Accept json
// @Produce json
// @Param  	file  formData file true "File"
// @Success 200 {object} Response{data=[]models.CreateCountry} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryUpload(c *gin.Context) {
	var countries = []models.CreateCountry{}
	file, err := c.FormFile("file")
	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, err.Error())
		return
	}
	dts := file.Filename

	err = c.SaveUploadedFile(file, dts)
	defer os.Remove(dts)

	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, err.Error())
		return
	}

	body, _ := os.ReadFile(dts)
	err = json.Unmarshal(body, &countries)
	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, err.Error())
		return
	}
	err = h.strg.Country().Upload(countries)
	if err != nil {
		handleResponse(c, http.StatusNotAcceptable, err.Error())
		return
	}
	handleResponse(c, http.StatusCreated, nil)
}
