CREATE TABLE IF NOT EXISTS patient(
    id SERIAL PRIMARY KEY,
    first_name   VARCHAR(50) NOT NULL,
	last_name    VARCHAR(50) NOT NULL,
	phone_number VARCHAR(15) NOT NULL CHECK (phone_number ~ '^\+54[0-9]{9,11}$'),
	dni         INT NOT NULL,
	health_insurance VARCHAR(50) NOT NULL,
	main       BOOLEAN NOT NULL,
    user_id INT8 NOT NULL,
    CONSTRAINT fk_patient_user_id FOREIGN KEY(user_id) REFERENCES "user"(id) ON DELETE CASCADE ON UPDATE CASCADE
);