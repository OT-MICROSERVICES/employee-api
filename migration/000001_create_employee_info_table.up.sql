CREATE TABLE IF NOT EXISTS employee_info (
    id text, name text, designation text, department text,
    joining_date date, address text, office_location text,
    status text, email text, phone_number text,
    PRIMARY KEY (id, joining_date)
) WITH CLUSTERING ORDER BY (joining_date DESC);