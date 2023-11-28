USE sample;

DROP TABLE IF EXISTS user;

CREATE TABLE user
(
    id       INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(40)
);

INSERT INTO user (name) VALUES ("佐藤太郎");
INSERT INTO user (name) VALUES ("鈴木一郎");
INSERT INTO user (name) VALUES ("users1");
INSERT INTO user (name) VALUES ("abc");
