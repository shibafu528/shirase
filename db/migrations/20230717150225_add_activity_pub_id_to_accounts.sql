-- migrate:up
ALTER TABLE accounts ADD COLUMN activity_pub_id TEXT;

-- migrate:down
ALTER TABLE accounts DROP COLUMN activity_pub_id;
