DROP TABLE IF EXISTS kafka;

CREATE TABLE kafka (
                       id SERIAL PRIMARY KEY,
                       producer_id int NULL,
                       producer_timestamp text NULL,
                       consumer_id int NULL,
                       consumer_timestamp text  NULL,
                       value text  NULL
);

SELECT COUNT(*) FROM KAFKA

-- BEGIN TRANSACTION;
-- INSERT INTO KAFKA (value) VALUES ('seed') RETURNING *;
-- COMMIT;
--
-- SELECT * FROM KAFKA
-- SELECT COUNT(*) FROM KAFKA