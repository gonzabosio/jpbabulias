CREATE TABLE IF NOT EXISTS "user"(
    id SERIAL PRIMARY KEY,
	email       VARCHAR(255) NOT NULL,
	password    VARCHAR(150) NOT NULL,
	birth_date TIMESTAMP NOT NULL,
	"admin"       BOOLEAN NOT NULL,
    CONSTRAINT unique_lower_email UNIQUE (LOWER(email))
);