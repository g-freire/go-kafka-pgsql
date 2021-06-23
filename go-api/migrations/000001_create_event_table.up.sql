DROP TABLE IF EXISTS kafka;

CREATE TABLE kafka (
                       id SERIAL PRIMARY KEY,
                       publisher_id int NULL,
                       consumer_id int NULL,
                       customer_name text NOT NULL,
                       value text NOT NULL
);

SELECT COUNT(*) FROM KAFKA

-- BEGIN TRANSACTION;
-- INSERT INTO KAFKA (value) VALUES ('seed') RETURNING *;
-- COMMIT;
--
-- SELECT * FROM KAFKA
-- SELECT COUNT(*) FROM KAFKA