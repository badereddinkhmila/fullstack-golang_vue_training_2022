CREATE TABLE IF NOT EXISTS "orders" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "status" varchar,
  "created_at" timestamp DEFAULT (now())
);


