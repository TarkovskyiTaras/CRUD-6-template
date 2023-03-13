DROP USER "crud-6";
CREATE USER "crud-6" WITH PASSWORD '12345' SUPERUSER;

DROP DATABASE "crud-6-db";
CREATE DATABASE "crud-6-db" OWNER "crud-6";

\connect crud-6-db

DROP TABLE persons;
CREATE TABLE persons (
    id INTEGER,
    first_name VARCHAR,
    last_name VARCHAR,
    dob  TIMESTAMP,
    home_address VARCHAR,
    cellphone VARCHAR
);

ALTER TABLE persons OWNER TO "crud-6";