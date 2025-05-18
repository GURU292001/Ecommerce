




-- customers definition

CREATE TABLE `customers` (
  `customer_id` varchar(100) NOT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `customer_email` varchar(100) DEFAULT NULL,
  `customer_address` varchar(200) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`customer_id`),
  UNIQUE KEY `customers_unique` (`customer_id`)
)

-- Orders definition

CREATE TABLE `Orders` (
  `order_id` varchar(25) NOT NULL,
  `product_id` varchar(50) DEFAULT NULL,
  `customer_id` varchar(50) DEFAULT NULL,
  `region_id` int(11) DEFAULT NULL,
  `date_of_sale` date DEFAULT NULL,
  `quantity_sold` int(11) DEFAULT NULL,
  `discount` decimal(10,2) DEFAULT NULL,
  `shipping_cost` int(11) DEFAULT NULL,
  `payment_method` varchar(50) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `updatedDate` datetime DEFAULT NULL,
  `id` int(11) DEFAULT NULL,
  PRIMARY KEY (`order_id`),
  KEY `Orders_customers_FK` (`customer_id`),
  KEY `Orders_products_FK` (`product_id`),
  KEY `Orders_region_FK` (`id`),
  CONSTRAINT `Orders_customers_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`),
  CONSTRAINT `Orders_products_FK` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`),
  CONSTRAINT `Orders_region_FK` FOREIGN KEY (`id`) REFERENCES `region` (`id`)
)

-- dev.products definition

CREATE TABLE `products` (
  `product_id` varchar(50) NOT NULL,
  `product_name` varchar(50) DEFAULT NULL,
  `unit_price` decimal(10,2) DEFAULT NULL,
  `product_category` varchar(50) DEFAULT NULL,
  `createdDate` datetime DEFAULT NULL,
  `UpdatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`product_id`),
  UNIQUE KEY `products_unique` (`product_id`)
) 

-- dev.region definition

CREATE TABLE `region` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `region_name` varchar(50) NOT NULL,
  `createdDate` datetime DEFAULT NULL,
  `updatedDate` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `region_unique` (`region_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;