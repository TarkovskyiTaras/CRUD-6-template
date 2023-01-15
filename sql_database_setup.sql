CREATE DATABASE persons;

CREATE TABLE persons (
    id INTEGER,
    first_name VARCHAR,
    last_name VARCHAR,
    dob  TIMESTAMP,
    home_address VARCHAR,
    cellphone VARCHAR
);

INSERT INTO persons
VALUES (1, 'Taras', 'Tarkovskyi', '2004-10-19 10:23:54+02', 'Kyiv', '0933115485');