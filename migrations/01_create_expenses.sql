CREATE TABLE IF NOT EXISTS expenses (
    id UUID PRIMARY KEY,
    amount DECIMAL NOT NULL CHECK (amount > 0),
    category TEXT NOT NULL,
    note TEXT,
    spent_on DATE NOT NULL,
    created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);