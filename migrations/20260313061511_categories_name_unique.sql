--20260313061511_categories_name_unique.sql
ALTER TABLE categories ADD CONSTRAINT categories_name_unique UNIQUE (name);
