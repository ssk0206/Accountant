CREATE TABLE IF NOT EXISTS bills (
    id INT AUTO_INCREMENT NOT NULL,
    roomid VARCHAR (3) UNIQUE NOT NULL,
    start_date VARCHAR(100),
    pre_meter_value FLOAT,
    new_meter_value FLOAT,
    bill FLOAT,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (roomid) REFERENCES students(id)
);