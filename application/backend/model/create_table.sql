CREATE
DATABASE miellink
CREATE TABLE users
(
    userID     INT PRIMARY KEY,
    username   VARCHAR(50) UNIQUE NOT NULL,
    `password` VARCHAR(50)        NOT NULL,
    realInfo   VARCHAR(100)
);
