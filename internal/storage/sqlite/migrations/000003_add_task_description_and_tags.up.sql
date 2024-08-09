
ALTER TABLE tasks ADD COLUMN description TEXT DEFAULT '';
ALTER TABLE tasks ADD COLUMN tags TEXT; -- tags will be a comma separated list of tags
