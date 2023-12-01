package storage

import "essy_travel/models"

type StorageI interface {
	City() CityRepoI
	Airport() AirportRepoI
	Country() CountryRepoI
}

type CountryRepoI interface {
	Create(req models.CreateCountry) (*models.Country, error)
	Update(req models.UpdateCountry) (*models.Country, error)
	GetById(req models.CountryPrimaryKey) (*models.Country, error)
	GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error)
	Delete(req models.CountryPrimaryKey) (string, error)
	Upload(req []models.CreateCountry) error
}

type CityRepoI interface {
	Create(req models.CreateCity) (*models.City, error)
	Update(req models.UpdateCity) (*models.City, error)
	GetById(req models.CityPrimaryKey) (*models.City, error)
	GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error)
	Delete(req models.CityPrimaryKey) (string, error)
	Upload(req []models.CreateCity) error
}

type AirportRepoI interface {
	Create(req models.CreateAirport) (*models.Airport, error)
	Update(req models.UpdateAirport) (*models.Airport, error)
	GetById(req models.AirportPrimaryKey) (*models.Airport, error)
	GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error)
	Delete(req models.AirportPrimaryKey) (string, error)
	Upload(req []models.CreateAirport) error
}
