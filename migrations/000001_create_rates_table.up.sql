CREATE TABLE IF NOT EXISTS exchange_rates (
    cur_id INT NOT NULL,
    date DATE NOT NULL,
    cur_abbreviation VARCHAR(10) NOT NULL,
    cur_name VARCHAR(100) NOT NULL,
    cur_scale INT NOT NULL,
    cur_official_rate FLOAT NOT NULL,
    PRIMARY KEY (cur_id, date),
    INDEX idx_date (date)
);
