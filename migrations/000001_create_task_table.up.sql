-- Description: Create the task table
CREATE TABLE IF NOT EXISTS tasks (
    id TEXT PRIMARY KEY, -- UUID
    title TEXT NOT NULL, -- Title of the task
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL
);