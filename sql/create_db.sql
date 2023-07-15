CREATE DATABASE IF NOT EXISTS knowledge;
USE knowledge;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS articles;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    admin boolean default false,
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE categories(
    id int auto_increment primary key,
    name varchar(50) not null,
    parent_id int null,
    FOREIGN KEY (parent_id) REFERENCES categories(id),
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE articles (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(1000) NOT NULL,
    image_url VARCHAR(1000),
    content LONGBLOB NOT NULL,
    user_id INT NOT NULL,
    category_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    createdAt timestamp default current_timestamp()
)ENGINE=INNODB;