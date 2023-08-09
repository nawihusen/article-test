CREATE TABLE `article` (
    `id`                    INT             NOT NULL AUTO_INCREMENT,
    `author`                TEXT            NOT NULL,
    `title`                 TEXT            NOT NULL,
    `body`                  TEXT            NOT NULL,
    `created`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
