DROP TABLE IF EXISTS kafka;

CREATE TABLE kafka (
	id SERIAL,
	value text
);
-- INSERT INTO kafka (value) VALUES ('seed') RETURNING *

-- BEGIN TRANSACTION;
--  INSERT INTO KAFKA (value) VALUES ('seed') RETURNING *;
-- COMMIT;

-- SELECT * FROM kafka