CREATE TABLE IF NOT EXISTS "app_users" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL ,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE app_users ADD COLUMN role varchar NOT NULL;