USE cursogoweb;
CREATE TABLE posts (
    id INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    body TEXT,
    created DATETIME DEFAULT NOW()
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO posts (title, body) VALUES ("Post test 1", "Content test 1");