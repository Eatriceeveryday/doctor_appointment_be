ALTER TABLE doctor_schedules ADD CONSTRAINT chk_schedules_day CHECK ( day IN ('Sunday' , 'Monday' , 'Tuesday' , 'Wednesday' , 'Thursday' , 'Friday' , 'Saturday') )