-- +migrate Up
CREATE TABLE users(
  id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  deleted_at datetime,
  UNIQUE KEY index_plans_on_name (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE users;
