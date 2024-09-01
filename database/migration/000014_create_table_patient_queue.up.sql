CREATE TABLE patient_queues (
    queue_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    prediction_time TIME NOT NULL ,
    queue_number INT NOT NULL,
    on_duty_id uuid NOT NULL ,
    doctor_id uuid NOT NULL ,
    patient_id uuid NOT NULL ,
    FOREIGN KEY (patient_id) REFERENCES patients (patient_id),
    FOREIGN KEY (doctor_id) REFERENCES doctors (doctor_id),
    FOREIGN KEY (on_duty_id) REFERENCES doctor_on_duty_schedules (on_duty_id)
);