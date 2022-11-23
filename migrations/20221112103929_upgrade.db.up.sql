CREATE DATABASE send_db;

USE send_db;

CREATE TABLE messages (
                          id INTEGER PRIMARY KEY AUTO_INCREMENT,
                          text text
);

CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTO_INCREMENT,
                       name varchar(255),
                       telegram_id INT,
                       first_name varchar(255),
                       last_name varchar(255),
);