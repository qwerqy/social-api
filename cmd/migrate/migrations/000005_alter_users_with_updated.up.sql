ALTER TABLE users
ADD COLUMN update_at timestamp(0) with time zone NOT NULL DEFAULT NOW();