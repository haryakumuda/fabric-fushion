CREATE DATABASE fabric_fushion;
USE DATABASE fabric_fushion;

-- table roles
create table roles (
    id int unsigned auto_increment primary key,
    role varchar(10)
);

-- table users
create table users (
    id int unsigned auto_increment primary key,
    role_id unsigned,
    email varchar(255) unique,
    password varchar(255) not null,
    foreign key (role_id) references roles(id)
);

-- table positions
create table positions (
    id int unsigned auto_increment primary key,
    position varchar(10)
);

-- table employees
create table employees (
    id int unsigned auto_increment primary key,
    position_id int unsigned,
    user_id int unsigned,
    name varchar(225),
    foreign key (position_id) references positions(id)
    foreign key (user_id) references users(id)
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
CREATE TABLE
  sales (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    customer_id INT UNSIGNED,
    order_date DATETIME NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
  );

-- table categories
create table categories (
    id int unsigned auto_increment primary key,
    category varchar(10)
);

-- Table Products
CREATE TABLE
  products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    category_id int unsigned,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    foreign key (category_id) references categories(id)
);

-- Table Products
CREATE TABLE sales_products (
    id int unsigned auto_increment primary key,
    sale_id int unsigned,
    product_id int unsigned,
    quantity int,
    foreign key (sale_id) references sales(id),
    foreign key (product_id) references products(id)
);

