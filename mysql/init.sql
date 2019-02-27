use test;

CREATE TABLE users (
  id int(10) unsigned not null auto_increment,
  name varchar(255),
  email varchar(255),
  primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO users (name, email) VALUES ('takahiro', 'takahiro@mail.com');
