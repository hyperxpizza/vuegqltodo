CREATE DATABASE vuegqltodo;

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    firstName VARCHAR(100) NOT NULL,
    lastName VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL
);

GRANT ALL PRIVILEGES ON TABLE contacts TO hyperxpizza;
GRANT USAGE, SELECT ON SEQUENCE contacts_id_seq TO hyperxpizza;