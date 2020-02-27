CREATE TABLE pages (
	title varchar(255) primary key not null,
    body text
);
INSERT INTO pages (title, body) VALUES('hello', 'world');