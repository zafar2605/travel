CREATE TABLE country(
  "guid"  UUID PRIMARY KEY,
  "title" VARCHAR(64),
  "code" VARCHAR(12),
  "continent" VARCHAR(12),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP
);
