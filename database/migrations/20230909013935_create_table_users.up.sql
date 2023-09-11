CREATE TABLE users(
    id VARCHAR(100) PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL ,
    email VARCHAR(100) UNIQUE NOT NULL ,
    password VARCHAR(255) NOT NULL ,
    is_deleted BOOLEAN DEFAULT false,
    created_at BIGINT NOT NULL ,
    updated_at BIGINT NOT NULL
)
