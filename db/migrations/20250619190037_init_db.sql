-- migrate:up

-- Enable uuid-ossp extension if not already enabled
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE incidents (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           title TEXT NOT NULL,
                           description TEXT NOT NULL,
                           affected_service TEXT NOT NULL,
                           ai_severity TEXT NOT NULL,
                           ai_category TEXT NOT NULL,
                           created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- migrate:down

DROP TABLE IF EXISTS incidents;
