CREATE TABLE doctors (
    doctor_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL ,
    image TEXT,
    hospital_id uuid NOT NULL
);