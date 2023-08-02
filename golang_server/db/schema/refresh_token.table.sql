CREATE TABLE IF NOT EXISTS "refresh_tokens" (
  "id" SERIAL PRIMARY KEY,
  "token" varchar,
  "expires_at" int,
  "user_id" int
);
