ALTER TABLE users
DROP COLUMN update_at;

ALTER TABLE users
ADD COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();

ALTER TABLE posts
DROP COLUMN update_at;

ALTER TABLE posts
ADD COLUMN updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();