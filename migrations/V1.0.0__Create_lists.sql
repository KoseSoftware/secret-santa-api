CREATE TABLE `lists` (
  `id`        CHAR(7)      NOT NULL,
  `organiser` VARCHAR(128) NOT NULL DEFAULT '',
  `email`     VARCHAR(128) NOT NULL DEFAULT '',
  `amount`    DOUBLE(8, 2) NOT NULL,
  `date`      DATETIME     NOT NULL,
  `location`  VARCHAR(255)          DEFAULT '',
  `notes`     VARCHAR(255)          DEFAULT '',
  `created`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `date_index` (`date`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;