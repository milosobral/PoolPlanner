-- Create the "pools" table
CREATE TABLE pools (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    href TEXT NOT NULL,
    address TEXT NOT NULL,
    neighborhood TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create the "events" table compatible with ICS format
CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    pool_id INT REFERENCES pools(id) ON DELETE CASCADE,
    summary TEXT NOT NULL, -- Event title or description
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    recurrence_rule TEXT, -- Optional: ICS-compatible recurrence rule
    recurrence_end_date TIMESTAMP, -- Optional: ICS-compatible recurrence end date
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Create the "calendars" table
CREATE TABLE calendars (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    pool_ids INT[] NOT NULL, -- Array of pool IDs
    unique_hash INT NOT NULL, -- Random unique hash
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

