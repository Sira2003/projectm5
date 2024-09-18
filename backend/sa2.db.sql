BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "genders" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"sex"	text,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "job_positions" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"job"	text,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "employees" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"first_name"	text,
	"last_name"	text,
	"birthday"	datetime,
	"address"	text,
	"tel"	text,
	"password"	text,
	"email"	text,
	"gender_id"	integer,
	"job_position_id"	integer,
	CONSTRAINT "fk_genders_employees" FOREIGN KEY("gender_id") REFERENCES "genders"("id"),
	CONSTRAINT "fk_job_positions_employees" FOREIGN KEY("job_position_id") REFERENCES "job_positions"("id"),
	CONSTRAINT "uni_employees_tel" UNIQUE("tel"),
	CONSTRAINT "uni_employees_email" UNIQUE("email"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "blood_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"blood_group"	text,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "patients" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"first_name"	text,
	"last_name"	text,
	"birthday"	datetime,
	"weight"	integer,
	"height"	integer,
	"drug_allergy"	text,
	"chronicdisease"	text,
	"tel"	text,
	"gender_id"	integer,
	"blood_type_id"	integer,
	CONSTRAINT "fk_blood_types_patients" FOREIGN KEY("blood_type_id") REFERENCES "blood_types"("id"),
	CONSTRAINT "fk_genders_patients" FOREIGN KEY("gender_id") REFERENCES "genders"("id"),
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "uni_patients_tel" UNIQUE("tel")
);
CREATE TABLE IF NOT EXISTS "dental_records" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"date"	datetime,
	"description"	text,
	"fees"	real,
	"installment"	real,
	"number_of_installments"	integer,
	"patient_id"	integer,
	"employee_id"	integer,
	CONSTRAINT "fk_patients_dental_records" FOREIGN KEY("patient_id") REFERENCES "patients"("id"),
	CONSTRAINT "fk_employees_dental_records" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "payments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"date"	datetime,
	"employee_id"	integer,
	CONSTRAINT "fk_employees_payments" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "requisitions" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"requisition_quantity"	integer,
	"time"	datetime,
	"note"	text,
	"employee_id"	integer,
	CONSTRAINT "fk_employees_requisitions" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "schedules" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"date"	datetime,
	"patient_id"	integer,
	CONSTRAINT "fk_patients_schedules" FOREIGN KEY("patient_id") REFERENCES "patients"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE INDEX IF NOT EXISTS "idx_genders_deleted_at" ON "genders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_job_positions_deleted_at" ON "job_positions" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_blood_types_deleted_at" ON "blood_types" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_patients_deleted_at" ON "patients" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_dental_records_deleted_at" ON "dental_records" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_payments_deleted_at" ON "payments" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_requisitions_deleted_at" ON "requisitions" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_schedules_deleted_at" ON "schedules" (
	"deleted_at"
);
COMMIT;
