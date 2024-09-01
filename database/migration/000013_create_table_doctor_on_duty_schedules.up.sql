CREATE TABLE doctor_on_duty_schedules(
    on_duty_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    day TEXT NOT NULL ,
    on_duty_start TIME NOT NULL ,
    on_duty_end TIME NOT NULL ,
    patient_limit INT NOT NULL ,
    doctor_id uuid NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctors (doctor_id),
    CHECK ( day IN ('Sunday' , 'Monday' , 'Tuesday' , 'Wednesday' , 'Thursday' , 'Friday' , 'Saturday'))
);