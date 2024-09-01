CREATE TABLE hospitals (
    hospital_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL ,
    address TEXT NOT NULL ,
    contact_number TEXT NOT NULL ,
    image TEXT
);