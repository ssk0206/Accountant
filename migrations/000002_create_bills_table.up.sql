CREATE TABLE IF NOT EXISTS bills (
    id INT AUTO_INCREMENT NOT NULL,
    room_id VARCHAR (3) UNIQUE,
    period VARCHAR(10),
    pre_meter_value FLOAT,
    new_meter_value FLOAT,
    bill FLOAT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (room_id) REFERENCES students(id)
    ON DELETE SET NULL
);