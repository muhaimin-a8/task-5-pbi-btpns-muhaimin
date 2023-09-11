CREATE TABLE photos(
    id VARCHAR(100) PRIMARY KEY ,
    title VARCHAR(50) NOT NULL ,
    caption VARCHAR(100) NOT NULL ,
    url VARCHAR(255) NOT NULL ,
    user_id VARCHAR(100) NOT NULL ,
    created_at BIGINT NOT NULL ,
    updated_at BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
)
