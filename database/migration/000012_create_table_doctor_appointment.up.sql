CREATE TABLE doctor_appointments(
    appointment_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    appointment_time TIMESTAMP NOT NULL ,
    patient_id uuid NOT NULL,
    schedule_id uuid NOT NULL ,
    doctor_id uuid NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients (patient_id),
    FOREIGN KEY (schedule_id) REFERENCES doctor_schedules (schedule_id),
    FOREIGN KEY (doctor_id) REFERENCES doctors (doctor_id),
    UNIQUE (schedule_id, appointment_time)
);