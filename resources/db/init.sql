CREATE TABLE IF NOT EXISTS metrics (
  id uuid NOT NULL PRIMARY KEY,
  name varchar NOT NULL,
  moisture integer NOT NULL,
  temp real NOT NULL,
  lux real NOT NULL,
  updated_at timestamp NOT NULL
);

CREATE INDEX IF NOT EXISTS metrics_name_idx ON metrics ("name");
CREATE INDEX IF NOT EXISTS metrics_updated_at_desc_idx ON metrics ("updated_at" DESC);

CREATE TABLE IF NOT EXISTS registrations (
  name varchar NOT NULL PRIMARY KEY,
  is_active boolean NOT NULL DEFAULT true,
  updated_at timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS alerts (
  id uuid NOT NULL PRIMARY KEY,
  name varchar NOT NULL,
  key_name integer NOT NULL,
  value real NOT NULL
);
CREATE INDEX IF NOT EXISTS alerts_name_idx ON alerts ("name");

CREATE TABLE IF NOT EXISTS config (
  name varchar UNIQUE NOT NULL,
  moisture_air integer NOT NULL,
  moisture_water integer NOT NULL,
  wait_time integer NOT NULL
);

CREATE INDEX IF NOT EXISTS config_name_idx ON config ("name");
