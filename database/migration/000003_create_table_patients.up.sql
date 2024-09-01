CREATE TABLE patients (
    patient_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL ,
    date_of_birth DATE NOT NULL ,
    gender TEXT NOT NULL,
    user_id uuid NOT NULL
);