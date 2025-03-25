-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS person (
   id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(), 
    monthly_income DECIMAL(10,2) NOT NULL,
    age INT NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    category VARCHAR(50) NOT NULL,
    balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

---- create above / drop below ----
DROP TABLE IF EXISTS person;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
