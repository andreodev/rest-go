CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY,
    name_product TEXT NOT NULL,
    price NUMERIC(12, 2) NOT NULL CHECK (price >= 0),
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
