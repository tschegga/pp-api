CREATE TABLE users (
    users_id INT AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(32) NOT NULL,
    role INT NOT NULL,
    PRIMARY KEY (users_id)
);

CREATE TABLE sessions (
    sessions_id INT AUTO_INCREMENT,
    athlete INT NOT NULL,
    start VARCHAR(255) NOT NULL,
    length INT NOT NULL,
    quality INT NOT NULL,
    PRIMARY KEY (sessions_id)
);

CREATE TABLE roles (
    roles_id INT AUTO_INCREMENT,
    name VARCHAR(255),
    PRIMARY KEY (roles_id)
);

INSERT INTO roles(name) VALUES ('admin');
INSERT INTO roles(name) VALUES ('athlete');
