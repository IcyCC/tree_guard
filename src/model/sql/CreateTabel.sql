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
);

CREATE TABLE operation(
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL UNIQUE ,
  INDEX idIndex(id),
  user_id INT DEFAULT -1,
  INDEX userIdIndex(user_id),
  tree_hash VARCHAR(255),
  operation VARCHAR(255) NULL,
  timestamp TIMESTAMP NOT NULL
);

SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
    op.timestamp as timestamp, op.user_id as user_id,
    u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE op.id = ? ) AS r ON r.tree_hash = t.hash;

SELECT r.id,r.operation,r.user_id,r.user_name,r.timestamp,t.hash,t.name,t.height
FROM tree AS t RIGHT JOIN
  (SELECT op.id as id,op.operation as operation,
          op.timestamp as timestamp, op.user_id as user_id,
          u.name as user_name, op.tree_hash as tree_hash
   FROM operation AS op JOIN user
     AS u on op.user_id=u.id WHERE u.id = ? ) AS r ON r.tree_hash = t.hash;