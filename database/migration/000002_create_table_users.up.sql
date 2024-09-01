CREATE TABLE users(
    user_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE ,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    contact_number TEXT NOT NULL
);