CREATE DATABASE IF NOT EXISTS knowledge;
USE knowledge;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    admin boolean default false,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;