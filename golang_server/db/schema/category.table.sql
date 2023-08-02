CREATE TABLE IF NOT EXISTS "categories" (
  id SERIAL PRIMARY KEY ,
  lable varchar UNIQUE ,
  description varchar,
  created_at timestamp DEFAULT (now()),
  created_by int
);
