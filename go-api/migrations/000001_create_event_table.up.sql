DROP TABLE IF EXISTS kafka;

CREATE TABLE kafka (
	id SERIAL PRIMARY KEY,
	value text
);
-- INSERT INTO kafka (value) VALUES ('test') RETURNING *
-- SELECT * FROM kafka
--
-- BEGIN TRANSACTION;
-- INSERT INTO KAFKA (value) VALUES ('tran test');
-- COMMIT;