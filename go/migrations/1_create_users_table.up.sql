CREATE TABLE IF NOT EXISTS users (
  id SERIAL primary key,
  name TEXT,
  email TEXT,
  password TEXT,
  token TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  created_by INTEGER,
  updated_by INTEGER,
  deleted_by INTEGER
)