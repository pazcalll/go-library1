DROP TABLE IF EXISTS borrows;

CREATE TABLE borrows (
    id int PRIMARY KEY AUTO_INCREMENT,
    user_id int NOT NULL,
    book_id int NOT NULL,
    is_return ENUM('YES', 'NO') NOT NULL,

    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);