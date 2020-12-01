BEGIN;
CREATE TYPE enum_live AS ENUM ('once', 'forever');
CREATE TYPE enum_status AS ENUM ('pending', 'done');
CREATE TABLE IF NOT EXISTS schedules (
    schedule_id SERIAL PRIMARY KEY,
    entry_id INT NOT NULL,
    name VARCHAR(30) NOT NULL,
    schedule VARCHAR(50) NOT NULL,
    action TEXT NOT NULL,
    live enum_live DEFAULT 'once',
    status enum_status DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
COMMIT;