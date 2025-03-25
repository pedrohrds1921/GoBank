-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS company (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(), 
    annual_revenue DECIMAL(15,2) NOT NULL,
    years_in_business INT NOT NULL,
    trade_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    corporate_email VARCHAR(255) UNIQUE NOT NULL,
    category VARCHAR(50) NOT NULL,
    balance DECIMAL(15,2) NOT NULL DEFAULT 0.00
);
---- create above / drop below ----
DROP TABLE IF EXISTS company;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
