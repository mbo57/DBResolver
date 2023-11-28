USE sample;

DROP TABLE IF EXISTS user;

CREATE TABLE user
(
    id       INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(40)
);

INSERT INTO user (name) VALUES ("佐藤太郎reader");
INSERT INTO user (name) VALUES ("鈴木一郎reader");
INSERT INTO user (name) VALUES ("user1reader");
INSERT INTO user (name) VALUES ("abcreader");
