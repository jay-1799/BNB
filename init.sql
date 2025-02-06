-- init.sql

-- Drop tables if they exist (for a clean start on container initialization)
DROP TABLE IF EXISTS bookings;
DROP TABLE IF EXISTS cabins;
DROP TABLE IF EXISTS guests;

CREATE TABLE guests (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    nationality VARCHAR(100),
    national_id VARCHAR(50),
    country_flag TEXT
);

CREATE TABLE cabins (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    max_capacity INTEGER NOT NULL,
    regular_price NUMERIC(10,2) NOT NULL,
    discount NUMERIC(10,2) NOT NULL DEFAULT 0,
    image TEXT,
    description TEXT
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    cabin_id INTEGER NOT NULL REFERENCES cabins(id),
    guest_id INTEGER NOT NULL REFERENCES guests(id),
    has_breakfast BOOLEAN DEFAULT false,
    observations TEXT,
    is_paid BOOLEAN DEFAULT false,
    num_guests INTEGER NOT NULL
);

INSERT INTO guests (full_name, email, nationality, national_id, country_flag)
VALUES 
('Jon Doe', 'hello@jon.io', 'Spain', '3525165164', 'https://flagcdn.com/pt.svg'),
('Jane Doe', 'jane@example.com', 'United States', '1234567890', 'https://flagcdn.com/us.svg'),
  ('John Smith', 'john@example.com', 'United Kingdom', '0987654321', 'https://flagcdn.com/gb.svg'),
  ('Alice Brown', 'alice@example.com', 'Canada', '1122334455', 'https://flagcdn.com/ca.svg'),
  ('Bob Martin', 'bob@example.com', 'Australia', '5566778899', 'https://flagcdn.com/au.svg');

INSERT INTO cabins (name, max_capacity, regular_price, discount, image, description)
VALUES 
  ('001', 2, 250.00, 0, 'https://your-cdn.com/cabin-001.jpg', 
    'Discover the ultimate luxury getaway for couples in cabin 001. Enjoy a cozy retreat with modern amenities.'),
  ('002', 4, 350.00, 10, 'https://your-cdn.com/cabin-002.jpg', 
    'Experience a family-friendly cabin in a scenic location with ample space and comfort.'),
  ('003', 3, 300.00, 5, 'https://your-cdn.com/cabin-003.jpg', 
    'Relax in a peaceful cabin with a blend of rustic charm and modern design, perfect for small groups.'),
  ('004', 2, 220.00, 0, 'https://your-cdn.com/cabin-004.jpg', 
    'A charming and intimate cabin ideal for a romantic getaway, featuring cozy interiors and scenic views.'),
  ('005', 5, 400.00, 15, 'https://your-cdn.com/cabin-005.jpg', 
    'Spacious cabin perfect for larger groups, offering a blend of luxury and convenience in a serene setting.');

INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '20 days', CURRENT_DATE, CURRENT_DATE + INTERVAL '7 days', 1, 1, TRUE, 
   'I have a gluten allergy and would like a gluten-free breakfast.', FALSE, 1);

-- Booking 2: Jane in cabin 002
INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '15 days', CURRENT_DATE + INTERVAL '2 days', CURRENT_DATE + INTERVAL '10 days', 2, 2, FALSE, 
   'No special requirements.', TRUE, 2);

-- Booking 3: John in cabin 003
INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '10 days', CURRENT_DATE + INTERVAL '5 days', CURRENT_DATE + INTERVAL '12 days', 3, 3, TRUE, 
   'Requesting an early check-in if possible.', FALSE, 3);

-- Booking 4: Alice in cabin 004
INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '5 days', CURRENT_DATE + INTERVAL '1 day', CURRENT_DATE + INTERVAL '8 days', 4, 4, TRUE, 
   'Would like a quiet room and a view of the mountains.', TRUE, 2);

-- Booking 5: Bob in cabin 005
INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '7 days', CURRENT_DATE + INTERVAL '3 days', CURRENT_DATE + INTERVAL '9 days', 5, 5, FALSE, 
   'Celebrating a birthday, please arrange a small cake if possible.', TRUE, 5);

-- Additional booking: Jonas books again in cabin 003
INSERT INTO bookings (created_at, start_date, end_date, cabin_id, guest_id, has_breakfast, observations, is_paid, num_guests)
VALUES 
  (CURRENT_DATE - INTERVAL '3 days', CURRENT_DATE + INTERVAL '10 days', CURRENT_DATE + INTERVAL '15 days', 3, 1, FALSE, 
   'Second booking for Jonas in a different cabin.', FALSE, 1);

