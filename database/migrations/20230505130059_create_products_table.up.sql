CREATE TABLE products(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_category_id` INT NULL,
    `title` VARCHAR(255) NOT NULL,
    `image` VARCHAR(255)NULL,
    `video` VARCHAR(255) NULL,
    `description` TEXT NULL,
    `price` INT NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_products_created_by (`created_by`),
    INDEX idx_products_updated_by (`updated_by`),
    CONSTRAINT FK_product_product_category_id FOREIGN KEY (`product_category_id`) REFERENCES product_categories(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_products_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_products_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)