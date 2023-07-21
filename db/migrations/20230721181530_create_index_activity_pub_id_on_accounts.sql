-- migrate:up
CREATE UNIQUE INDEX accounts_activity_pub_id ON accounts(activity_pub_id);

-- migrate:down
DROP INDEX accounts_activity_pub_id;
