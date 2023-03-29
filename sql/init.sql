CREATE TABLE IF NOT EXISTS todos (
  id serial PRIMARY KEY,
  name varchar(50) UNIQUE NOT NULL,
  completed boolean NOT NULL,
  due timestamp
);


