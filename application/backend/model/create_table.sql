CREATE
DATABASE fabrictrace
CREATE TABLE users
(
    user_id    INT PRIMARY KEY,
    username   VARCHAR(50) UNIQUE NOT NULL,
    `password` VARCHAR(50)        NOT NULL,
    realInfo   VARCHAR(100)
);
