DROP TABLE IF EXISTS "order_item";
CREATE TABLE IF NOT EXISTS "order_item" (
  "id" SERIAL,
  "quantity" int DEFAULT 1,
  "order_id" int,
  "product_id" int,
  PRIMARY KEY(order_id,product_id)
);