CREATE DATABASE railway;
USE railway;

-- table roles
CREATE TABLE roles (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    role VARCHAR(10)
);

-- table users
CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    role_id INT UNSIGNED,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255) NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- table positions
CREATE TABLE positions (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    position VARCHAR(10)
);

-- table employees
CREATE TABLE employees (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    position_id INT UNSIGNED,
    user_id INT UNSIGNED,
    name VARCHAR(225),
    FOREIGN KEY (position_id) REFERENCES positions(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table Customers
CREATE TABLE customers (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED,
    phone_number VARCHAR(30) NOT NULL,
    name VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table Sales
CREATE TABLE sales (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    customer_id INT UNSIGNED,
    order_date DATETIME NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- table categories
CREATE TABLE categories (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    category VARCHAR(10)
);

-- Table Products
CREATE TABLE products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    category_id INT UNSIGNED,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock int not null,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Table sales_products
CREATE TABLE sales_products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sale_id INT UNSIGNED,
    product_id INT UNSIGNED,
    quantity INT,
    FOREIGN KEY (sale_id) REFERENCES sales(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);