CREATE TABLE tree(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE,
  INDEX idIndex (id),
  name VARCHAR(256)  DEFAULT '',
  INDEX nameIndex (name),
  hash VARCHAR(256)  UNIQUE DEFAULT '',
  INDEX hashIndex(hash),
  height INT DEFAULT 0
)

