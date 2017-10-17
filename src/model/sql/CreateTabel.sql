CREATE TABLE tree(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
  INDEX idIndex (id),
  name VARCHAR(255)  DEFAULT '',
  INDEX nameIndex (name),
  hash VARCHAR(255)  UNIQUE DEFAULT '',
  INDEX hashIndex(hash),
  height FLOAT DEFAULT 0
);

CREATE TABLE user(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE ,
  INDEX idIndex (id),
  name VARCHAR(255) DEFAULT '',
  INDEX nameIndex(name),
  password VARCHAR(255) DEFAULT ''
)

