CREATE TABLE doctor_schedules(
    schedule_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    day TEXT,
    hour TIME,
    doctor_id uuid NOT NULL,
    UNIQUE (day, hour, doctor_id)
);
