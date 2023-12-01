package api

import (
	"essy_travel/api/handler"
	"essy_travel/config"
	"essy_travel/storage"

	"github.com/gin-gonic/gin"

	_ "essy_travel/api/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	// City ...
	r.POST("/city", handler.CreateCity)
	r.GET("/city/:id", handler.CityGetById)
	r.GET("/city", handler.CityGetList)
	r.PUT("/city", handler.CityUpdate)
	r.DELETE("/city", handler.CityDelete)
	r.POST("/city/:upload", handler.CityUpload)

	// Country
	r.POST("/country", handler.CreateCountry)
	r.GET("/country/:id", handler.CountryGetById)
	r.GET("/country", handler.CountryGetList)
	r.PUT("/country", handler.CountryUpdate)
	r.DELETE("/country", handler.CountryDelete)
	r.POST("/country/:upload", handler.CountryUpload)

	// Airport
	r.POST("/airport", handler.CreateAirport)
	r.GET("/airport/:id", handler.AirportGetById)
	r.GET("/airport", handler.AirportGetList)
	r.PUT("/airport", handler.AirportUpdate)
	r.DELETE("/airport", handler.AirportDelete)
	r.POST("/airport/:upload", handler.AirportUpload)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
