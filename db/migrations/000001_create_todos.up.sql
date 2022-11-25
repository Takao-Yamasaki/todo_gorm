CREATE TABLE `todos` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `text`     VARCHAR(191) NOT NULL,
	`status`   INT(11) NOT NULL,
	`deadline` INT(11) NOT NULL,
    PRIMARY KEY(`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;