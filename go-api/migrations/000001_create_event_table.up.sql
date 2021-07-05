DROP TABLE IF EXISTS events;

CREATE TABLE events (
                       id SERIAL PRIMARY KEY,
                       producer_id int NULL,
                       producer_timestamp text NULL,
                       consumer_id int NULL,
                       consumer_timestamp text  NULL,
                       value int  NULL
);

SELECT COUNT(*) FROM events

-- BEGIN TRANSACTION;
-- INSERT INTO events (value) VALUES ('seed') RETURNING *;
-- COMMIT;
--
SELECT COUNT(*) FROM events
SELECT * FROM events
SELECT * FROM events WHERE producer_id = 1
SELECT * FROM events WHERE producer_id = 2