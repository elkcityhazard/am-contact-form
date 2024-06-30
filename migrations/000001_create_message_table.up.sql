CREATE TABLE Message (
    message_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NOW(),
    version INT DEFAULT 1,
    name VARCHAR(255) DEFAULT "",
    email VARCHAR(255) NOT NULL UNIQUE,
    message_content TEXT,
    ip_address VARCHAR(15)
);
