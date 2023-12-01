package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type CountryRepo struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (c *CountryRepo) Create(req models.CreateCountry) (*models.Country, error) {
	query := `
	INSERT INTO country(
		guid,
		title,
		code,
		continent,
		updated_at)
		VALUES ($1,$2,$3,4$,NOW())`
	guid := uuid.New().String()
	_, err := c.db.Exec(query, guid, req.Title, req.Code, req.Continent)
	if err != nil {
		return &models.Country{}, err
	}
	return c.GetById(models.CountryPrimaryKey{Guid: guid})
}

func (c *CountryRepo) GetById(req models.CountryPrimaryKey) (*models.Country, error) {
	var country = models.Country{}
	query := `
		SELECT 
			guid,
			title,
			code,
			continent,
			created_at,
			updated_at 
		FROM country
	`
	resp := c.db.QueryRow(query, req.Guid)
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	err := resp.Scan(
		&country.Guid,
		&country.Title,
		&country.Code,
		&country.Code,
		&country.CreatedAt,
		&country.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (c *CountryRepo) GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error) {
	var (
		resp   = models.GetListCountryResponse{}
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
			"code",
			"continent",
			"created_at",
			"updated_at"
		FROM country
	`
	query += where + limit + offset

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Guid      sql.NullString
			Title     sql.NullString
			Code      sql.NullString
			Continent sql.NullString
			CreatedAt sql.NullString
			UpdatedAt sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Guid,
			&Title,
			&Code,
			&Continent,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Countries = append(resp.Countries, models.Country{
			Guid:      Guid.String,
			Title:     Title.String,
			Code:      Code.String,
			Continent: Continent.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		})
	}

	return &resp, nil
}

func (c *CountryRepo) Update(req models.UpdateCountry) (*models.Country, error) {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>______", req)
	query := `
		UPDATE country SET 
			"title" = $1,
			"code" = $2,
			"continent" = $3,
			"updated_at" = NOW() 
		WHERE 
			guid = $4`
	_, err := c.db.Exec(query, req.Title, req.Code, req.Continent, req.Guid)
	if err != nil {
		return &models.Country{}, err
	}
	fmt.Println("okokokokokokkkooko")

	return &models.Country{}, nil
}

func (c *CountryRepo) Delete(req models.CountryPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM country WHERE guid = $1`, req.Guid)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CountryRepo) Upload(req []models.CreateCountry) error {
	query := `
		INSERT INTO country(
			"guid",
			"title",
			"code",
			"continent",
			"updated_at") VALUES
			($1, $2, $3, $4, NOW())
	`
	for _, v := range req {
		
		guid := uuid.New().String()
		_, err := c.db.Exec(query, guid, v.Title, v.Code, v.Continent)
		if err != nil {
			return err
		}
	}

	return nil
}
