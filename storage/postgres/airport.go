package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type AirportRepo struct {
	db *sql.DB
}

func NewAirportRepo(db *sql.DB) *AirportRepo {
	return &AirportRepo{
		db: db,
	}
}

func (a *AirportRepo) Create(req models.CreateAirport) (*models.Airport, error) {
	query := `INSERT INTO aiport(guid,title,country_id,city_id,longitude,radius,image,adress,timezone_id,country,city,search_text,code,product_count,gmt,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,NOW())`
	guid := uuid.New().String()
	_, err := a.db.Exec(query,
		guid,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Adress,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}
	return a.GetById(models.AirportPrimaryKey{Guid: guid})
}

func (a *AirportRepo) GetById(req models.AirportPrimaryKey) (*models.Airport, error) {
	query := `
		SELECT
			"title"       
			"country_id"   
			"city_id"      
			"latitude"    
			"longitude"   
			"radius"      
			"image"       
			"adress"      
			"timezone_id"  
			"country"     
			"city"        
			"search_text"  
			"code"        
			"product_count"
			"gmt"         
			"created_at"  
			"updated_at"   
		FROM airport
		WHERE guid = $1
	`

	var (
		Guid         sql.NullString
		Title        sql.NullString
		CountryId    sql.NullString
		CityId       sql.NullString
		Latitude     sql.NullFloat64
		Longitude    sql.NullFloat64
		Radius       sql.NullFloat64
		Image        sql.NullString
		Adress       sql.NullString
		TimezoneId   sql.NullString
		Country      sql.NullString
		City         sql.NullString
		SearchText   sql.NullString
		Code         sql.NullString
		ProductCount sql.NullInt64
		Gmt          sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString
	)

	err := a.db.QueryRow(query, req.Guid).Scan(
		&Guid,
		&Title,
		&CountryId,
		&CityId,
		&Latitude,
		&Longitude,
		&Radius,
		&Image,
		&Adress,
		&TimezoneId,
		&Country,
		&City,
		&SearchText,
		&Code,
		&ProductCount,
		&Gmt,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &models.Airport{
		Guid:         Guid.String,
		Title:        Title.String,
		CountryId:    CountryId.String,
		CityId:       CityId.String,
		Latitude:     Latitude.Float64,
		Longitude:    Longitude.Float64,
		Radius:       Radius.Float64,
		Image:        Image.String,
		Adress:       Adress.String,
		TimezoneId:   TimezoneId.String,
		Country:      Country.String,
		City:         City.String,
		SearchText:   SearchText.String,
		Code:         Code.String,
		ProductCount: int(ProductCount.Int64),
		Gmt:          Gmt.String,
		CreatedAt:    CreatedAt.String,
		UpdatedAt:    UpdatedAt.String,
	}, nil
}

func (a *AirportRepo) GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error) {
	var (
		resp   = models.GetListAirportResponse{}
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
			"city_id",
			"latitude",
			"longitude",
			"radius",
			"image",
			"adress",
			"timezone_id",
			"country",
			"city",
			"search_text",
			"code",
			"product_count",
			"gmt",
			"created_at",
			"updated_at"
		FROM airport
	`
	query += where + limit + offset

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Guid         sql.NullString
			Title        sql.NullString
			CountryId    sql.NullString
			CityId       sql.NullString
			Latitude     sql.NullFloat64
			Longitude    sql.NullFloat64
			Radius       sql.NullFloat64
			Image        sql.NullString
			Adress       sql.NullString
			TimezoneId   sql.NullString
			Country      sql.NullString
			City         sql.NullString
			SearchText   sql.NullString
			Code         sql.NullString
			ProductCount sql.NullInt64
			Gmt          sql.NullString
			CreatedAt    sql.NullString
			UpdatedAt    sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Guid,
			&Title,
			&CountryId,
			&CityId,
			&Latitude,
			&Longitude,
			&Radius,
			&Image,
			&Adress,
			&TimezoneId,
			&Country,
			&City,
			&SearchText,
			&Code,
			&ProductCount,
			&Gmt,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Airports = append(resp.Airports, models.Airport{
			Guid:         Guid.String,
			Title:        Title.String,
			CountryId:    CountryId.String,
			CityId:       CityId.String,
			Latitude:     Latitude.Float64,
			Longitude:    Longitude.Float64,
			Radius:       Radius.Float64,
			Image:        Image.String,
			Adress:       Adress.String,
			TimezoneId:   TimezoneId.String,
			Country:      Country.String,
			City:         City.String,
			SearchText:   SearchText.String,
			Code:         Code.String,
			ProductCount: int(ProductCount.Int64),
			Gmt:          Gmt.String,
			CreatedAt:    CreatedAt.String,
			UpdatedAt:    UpdatedAt.String,
		})
	}

	return &resp, nil
}

func (a *AirportRepo) Update(req models.UpdateAirport) (*models.Airport, error) {

	query := `UPDATE airports SET guid=$1,title=$2,country_id=$3,city_id=$4,longitude=$5,radius=$6,image=$7,adress=$8,timezone_id=$9,country=$10,city=$11,search_text=$12,code=$13,product_count=$14,gmt=$15,updated_at = NOW() WHERE id = $16`
	_, err := a.db.Exec(
		query,
		req.Guid,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Adress,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}

	return a.GetById(models.AirportPrimaryKey{Guid: req.Guid})
}

func (a *AirportRepo) Delete(req models.AirportPrimaryKey) (string, error) {

	_, err := a.db.Exec(`DELETE FROM airports WHERE id = $1`, req.Guid)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *AirportRepo) Upload(req []models.CreateAirport) error {
	query := `
		INSERT INTO airport(
			"guid",
			"title",
			"country_id",
			"city_id",
			"latitude",
			"longitude",
			"radius",
			"image",
			"adress",
			"timezone_id",
			"country",
			"city",
			"search_text",
			"code",
			"product_count",
			"gmt",
			"updated_at") VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, NOW())
	`
	for _, v := range req {
		guid := uuid.New().String()

		_, err := c.db.Exec(query, guid, v.Title, v.CountryId, v.CityId, v.Latitude, v.Longitude,
			v.Radius, v.Image, v.Adress, v.TimezoneId, v.Country, v.City, v.SearchText, v.Code,
			v.ProductCount, v.Gmt)
		if err != nil {
			return err
		}
	}

	return nil
}
