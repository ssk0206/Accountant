CREATE TABLE IF NOT EXISTS students(
    id INT (3) UNIQUE NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id)
);