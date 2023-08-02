CREATE TABLE IF NOT EXISTS "products" (
  "id" SERIAL PRIMARY KEY,
  "name_product" varchar,
  "description_product" varchar,
  "price_product" float,
  "quantity" int,
  "is_available" boolean,
  "image_product" bytea,
  "created_at" timestamp DEFAULT (now()),
  "category_id" int,
  "created_by" int
);

CREATE INDEX "product name" ON "products" USING BTREE("name_product");

CREATE INDEX "product description" ON "products" USING BTREE("description_product");
