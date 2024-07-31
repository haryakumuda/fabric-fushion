-- Insert roles
INSERT INTO roles (role) VALUES 
('admin'), 
('customer');

-- Insert users
INSERT INTO users (role_id, email, password) VALUES
(1, 'a@gmail.com', 'pass1'),
(2, 'b@gmail.com', 'pass2'),
(2, 'c@gmail.com', 'pass3'),
(2, 'd@gmail.com', 'pass4'),
(1, 'e@gmail.com', 'pass5'),
(2, 'f@gmail.com', 'pass6'),
(2, 'g@gmail.com', 'pass7'),
(1, 'h@gmail.com', 'pass8'),
(2, 'i@gmail.com', 'pass9'),
(1, 'j@gmail.com', 'pass10');

-- Insert positions
INSERT INTO positions (position) VALUES 
('Manager'), 
('Sales'),
('Intern');

-- Insert employees
INSERT INTO employees (position_id, user_id, name) VALUES
(1, 1, 'Alice Johnson'),
(2, 2, 'Bob Smith'),
(3, 3, 'Carol White'),
(1, 4, 'David Brown'),
(2, 5, 'Eve Black');

-- Insert customers
INSERT INTO customers (user_id, phone_number, name) VALUES
(7, '789-012-3456', 'Hillary Clinton'),
(8, '890-123-4567', 'Isaac Newton'),
(6, '210-987-6543', 'Jack Daniels'),
(9, '345-678-9012', 'Karen Gillan'),
(10, '567-890-1234', 'Leonardo DiCaprio');

-- Insert categories
INSERT INTO categories (category) VALUES
('Tops'), 
('Bottoms'), 
('Footwear'), 
('Accs'), 
('Outerwear');

-- Insert products
INSERT INTO products (category_id, name, price) VALUES
(1, 'T-Shirt', 19.99),
(1, 'Blouse', 29.99),
(2, 'Jeans', 49.99),
(2, 'Shorts', 25.99),
(3, 'Sneakers', 79.99),
(3, 'Boots', 129.99),
(4, 'Watch', 99.99),
(4, 'Hat', 15.99),
(5, 'Jacket', 89.99),
(5, 'Coat', 129.99);
-- Insert sales
INSERT INTO sales (customer_id, order_date) VALUES
(1, '2024-07-30 10:00:00'),
(2, '2024-07-30 11:00:00'),
(3, '2024-07-31 09:00:00'),
(4, '2024-07-31 14:00:00'),
(5, '2024-08-01 10:30:00');

-- Insert sales_products
INSERT INTO sales_products (sale_id, product_id, quantity) VALUES
(1, 1, 2),
(2, 2, 1),
(3, 3, 3),
(4, 4, 2),
(5, 5, 1);

