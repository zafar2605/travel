package models

type City struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateCity struct {
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
}

type UpdateCity struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	CountryId   string `json:"country_id"`
	CityCode    string `json:"city_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Offset      string `json:"offset"`
	TimezoneId  string `json:"timezone_id"`
	CountryName string `json:"country_name"`
}

type CityPrimaryKey struct {
	Guid string `json:"guid"`
}

type GetListCityRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListCityResponse struct {
	Count  int    `json:"count"`
	Cities []City `json:"cities"`
}
