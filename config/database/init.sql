CREATE DATABASE payment_app;

CREATE TABLE user_credential (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    role VARCHAR(50) NOT NULL
);

-- dummy user with uuid and password = admin with hashed password
INSERT INTO user_credential (id, username, password, role) VALUES (
    'f1b9a7a0-5f1a-11eb-ae93-0242ac130002',
    'admin',
    '$2a$14$Ty0uDVUu93MI5YrEkjiAPuu.n5G4vflQrUGYIsl8L08SnAfxKTRBe',
    'admin'
);

CREATE TABLE profile (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(100) NOT NULL,
    balance VARCHAR (10000000)
);

CREATE TABLE merchant (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    name_merchant VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    phone VARCHAR(100) NOT NULL,
    balance VARCHAR (10000000)
);

CREATE TABLE transfer (
    id VARCHAR(100) NOT NULL PRIMARY KEY,
    sender_id VARCHAR(100) NOT NULL,
    receiver_id VARCHAR(100) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    description VARCHAR(255)

    FOREIGN KEY (sender_id) REFERENCES profile(id),
    FOREIGN KEY (receiver_id) REFERENCES merchant(id)
);











