CREATE TABLE city(
  "guid" UUID PRIMARY KEY,
  "title" VARCHAR(48),
  "country_id" UUID,
  "city_code" VARCHAR(120),
  "latitude" VARCHAR(120),
  "longitude" VARCHAR(120),
  "offset" VARCHAR(120),
  "timezone_id" UUID,
  "country_name" VARCHAR(128),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);




jq -c '.[]' city.json >> your_data.json


\copy temp(data) from './mock_data/your_data.json';


INSERT INTO city (
  "title",
  "country_id",
  "city_code",
  "latitude",
  "longitude",
  "offset",
  "timezone_id",
  "country_name",
  "updated_at"
)
SELECT
  data ->> 'title',
  CAST(data ->> 'country_id' AS UUID),
  data ->> 'city_code',
  data ->> 'latitude',
  data ->> 'longitude',
  data ->> 'offset',
  CAST(data ->> 'timezone_id' AS UUID),
  data ->> 'country_name',
  now()
FROM
  temp;


DELETE FROM temp WHERE length(data ->> 'country_id') = 2;

