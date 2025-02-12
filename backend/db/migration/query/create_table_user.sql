CREATE TABLE IF NOT EXISTS "user"(
    id SERIAL PRIMARY KEY,
    first_name   VARCHAR(50) NOT NULL,
	last_name    VARCHAR(50) NOT NULL,
	email       VARCHAR(255) NOT NULL,
	password    VARCHAR(150) NOT NULL,
	phone_number VARCHAR(15) NOT NULL CHECK (phone_number ~ '^\+54[0-9]{9,11}$'),
	dni         INT NOT NULL,
	"admin"       BOOLEAN NOT NULL,
	health_insurance_id INT8 NOT NULL,
    CONSTRAINT unique_lower_email UNIQUE (LOWER(email)),
	CONSTRAINT fk_user_insurance FOREIGN KEY(health_insurance_id) REFERENCES health_insurance(id) ON UPDATE CASCADE
);