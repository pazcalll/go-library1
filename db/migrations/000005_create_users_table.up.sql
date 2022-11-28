DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(80) NOT NULL,
    img_url varchar(255) NOT NULL
);