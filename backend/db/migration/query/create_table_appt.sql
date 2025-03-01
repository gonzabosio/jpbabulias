CREATE TABLE IF NOT EXISTS appointment (
	id SERIAL PRIMARY KEY,
	appt_date TIMESTAMP NOT NULL,
	subject VARCHAR(50) NOT NULL,
	patient_id   INT8 NOT NULL,
	CONSTRAINT fk_appt_patient FOREIGN KEY(patient_id) REFERENCES patient(id) ON UPDATE CASCADE ON DELETE CASCADE
);