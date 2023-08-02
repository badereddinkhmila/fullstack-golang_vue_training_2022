ALTER TABLE "order_item" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");