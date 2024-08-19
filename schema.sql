
CREATE TABLE file (
	id TEXT PRIMARY KEY NOT NULL
);

CREATE TABLE attribute (
	file NOT NULL REFERENCES file,
	scope TEXT NOT NULL,
	name TEXT NOT NULL,
	UNIQUE (file, scope, name)
);

CREATE TABLE attribute_value (
	attribute NOT NULL REFERENCES attribute,
	value TEXT NOT NULL
);
