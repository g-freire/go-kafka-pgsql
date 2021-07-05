DROP TABLE IF EXISTS processed_message;

CREATE TABLE processed_message (
                       msg_id SERIAL PRIMARY KEY
);

SELECT COUNT(*) FROM processed_message
SELECT * FROM processed_message
INSERT INTO processed_message(msg_id) VALUES (123)