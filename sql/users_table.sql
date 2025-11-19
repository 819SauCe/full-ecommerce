CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_img VARCHAR(250),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_area_code VARCHAR(4),
    country_code VARCHAR(4),
    phone VARCHAR(16),
    email VARCHAR(200) UNIQUE NOT NULL,
    password VARCHAR(300) NOT NULL,
    role VARCHAR(25) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);