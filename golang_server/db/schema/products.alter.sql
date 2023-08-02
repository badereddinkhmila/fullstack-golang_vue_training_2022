ALTER TABLE "products" ADD FOREIGN KEY ("created_by") REFERENCES "app_users" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

CREATE INDEX "product_status" ON "products" USING BTREE("created_by", "is_available");

CREATE INDEX "provider" ON "products" USING BTREE("created_by");
