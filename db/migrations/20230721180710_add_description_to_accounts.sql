-- migrate:up
ALTER TABLE accounts ADD COLUMN description TEXT;

-- migrate:down
ALTER TABLE accounts DROP COLUMN description;
