
-- Users table
CREATE TABLE users (
    -- Core fields
    id UUID DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    deleted_by VARCHAR(255)
);

-- Indexes
CREATE INDEX idx_users_email ON users(email);