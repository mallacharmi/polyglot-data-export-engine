#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE TABLE IF NOT EXISTS public.records (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    value DECIMAL(18,4) NOT NULL,
    metadata JSONB NOT NULL
);

DO \$\$
BEGIN
   IF (SELECT COUNT(*) FROM records) = 0 THEN
      INSERT INTO records (name, value, metadata)
      SELECT
          'User_' || gs,
          random() * 10000,
          jsonb_build_object(
              'category', 'cat_' || (gs % 10),
              'details', jsonb_build_object(
                  'score', random(),
                  'active', (gs % 2 = 0)
              )
          )
      FROM generate_series(1, 10000000) AS gs;
   END IF;
END
\$\$;

EOSQL