package postgres

import (
	"database/sql"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"fmt"

	"github.com/google/uuid"
)

type CityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (c *CityRepo) Create(req models.CreateCity) (*models.City, error) {

	query := `
		INSERT INTO city(
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"updated_at"
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())`

	id := uuid.New().String()
	_, err := c.db.Exec(query,
		id,
		req.Title,
		helpers.NewNullString(req.CountryId),
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		helpers.NewNullString(req.TimezoneId),
		req.CountryName,
	)
	if err != nil {
		return &models.City{}, err
	}

	return c.GetById(models.CityPrimaryKey{Guid: id})
}

func (c *CityRepo) GetById(req models.CityPrimaryKey) (*models.City, error) {
	query := `
		SELECT
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
		WHERE guid = $1
	`

	var (
		Title       sql.NullString
		CountryId   sql.NullString
		CityCode    sql.NullString
		Latitude    sql.NullString
		Longitude   sql.NullString
		Offset      sql.NullString
		TimezoneId  sql.NullString
		CountryName sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	err := c.db.QueryRow(query, req.Guid).Scan(
		&req.Guid,
		&Title,
		&CountryId,
		&CityCode,
		&Latitude,
		&Longitude,
		&Offset,
		&TimezoneId,
		&CountryName,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.City{
		Guid:        req.Guid,
		Title:       Title.String,
		CountryId:   CountryId.String,
		CityCode:    CityCode.String,
		Latitude:    Latitude.String,
		Longitude:   Longitude.String,
		Offset:      Offset.String,
		TimezoneId:  TimezoneId.String,
		CountryName: CountryName.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}, nil
}

func (c *CityRepo) GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		resp   = models.GetListCityResponse{}
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"created_at",
			"updated_at"
		FROM city
	`
	query += where + limit + offset

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Guid        sql.NullString
			Title       sql.NullString
			CountryId   sql.NullString
			CityCode    sql.NullString
			Latitude    sql.NullString
			Longitude   sql.NullString
			Offset      sql.NullString
			TimezoneId  sql.NullString
			CountryName sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Guid,
			&Title,
			&CountryId,
			&CityCode,
			&Latitude,
			&Longitude,
			&Offset,
			&TimezoneId,
			&CountryName,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Cities = append(resp.Cities, models.City{
			Guid:        Guid.String,
			Title:       Title.String,
			CountryId:   CountryId.String,
			CityCode:    CityCode.String,
			Latitude:    Latitude.String,
			Longitude:   Longitude.String,
			Offset:      Offset.String,
			TimezoneId:  TimezoneId.String,
			CountryName: CountryName.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &resp, nil
}

func (c *CityRepo) Update(req models.UpdateCity) (*models.City, error) {
	query := `
		UPDATE city SET 
			"title" = $1,
			"country_id" = $2,
			"city_code" = $3,
			"latitude" = $4,
			"longitude" = $5,
			"offset" = $6,
			"timezone_id" = $7,
			"country_name" = $8
		WHERE "guid" = $9
	`
	_, err := c.db.Exec(
		query,
		req.Title,
		req.CountryId,
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		req.TimezoneId,
		req.CountryName,
		req.Guid,
	)
	if err != nil {
		return &models.City{}, err
	}

	return c.GetById(models.CityPrimaryKey{Guid: req.Guid})
}

func (c *CityRepo) Delete(req models.CityPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM city WHERE guid = $1`, req.Guid)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CityRepo) Upload(req []models.CreateCity) error {
	query := `
		INSERT INTO city(
			"guid",
			"title",
			"country_id",
			"city_code",
			"latitude",
			"longitude",
			"offset",
			"timezone_id",
			"country_name",
			"updated_at") VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`
	for _, v := range req {
		
		guid := uuid.New().String()
		_, err := c.db.Exec(query, guid, v.Title, v.CountryId, v.CityCode, v.Latitude,
			v.Longitude, v.Offset, v.TimezoneId, v.CountryName)
		if err != nil {
			return err
		}
	}

	return nil
}
