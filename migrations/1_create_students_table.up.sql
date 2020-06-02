CREATE TABLE IF NOT EXISTS accountant.students(
    id INT AUTO_INCREMENT NOT NULL,
    room_id VARCHAR (3) NOT NULL,
    name VARCHAR (50) NOT NULL,
    pre_meter_value FLOAT,
    new_meter_value FLOAT,
    bill FLOAT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id, room_id)
);