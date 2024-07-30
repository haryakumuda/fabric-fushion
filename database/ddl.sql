CREATE DATABASE fabric_fushion;

USE DATABASE fabric_fushion;

-- Table Users
CREATE TABLE
  users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
  );

-- Table Employees
CREATE TABLE
  employees (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED,
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
  );

-- Table Customers
CREATE TABLE
  customers (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED,
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(30) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
  );

-- Table Products
CREATE TABLE
  products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100) NOT NULL
  );

-- Table Sales
CREATE TABLE
  sales (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_date DATETIME NOT NULL,
    customer_id INT UNSIGNED,
    FOREIGN KEY (customer_id) REFERENCES customers (id)
  );