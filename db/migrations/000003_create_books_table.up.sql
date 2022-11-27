DROP TABLE IF EXISTS books;

CREATE TABLE books (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(80) NOT NULL,
    author varchar(81) NOT NULL,
    img_url varchar(255) NOT NULL
);