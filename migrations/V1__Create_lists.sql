CREATE TABLE `lists` (
  `id`        INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `organiser` VARCHAR(128)     NOT NULL DEFAULT '',
  `email`     VARCHAR(128)     NOT NULL DEFAULT '',
  `amount`    DOUBLE(8, 2)     NOT NULL,
  `date`      DATETIME         NOT NULL,
  `location`  VARCHAR(255)              DEFAULT '',
  `notes`     VARCHAR(255)              DEFAULT '',
  `created`   DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated`   DATETIME         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 2304
  DEFAULT CHARSET = latin1;

INSERT INTO `lists` (
  `id`,
  `organiser`,
  `email`,
  `amount`,
  `date`,
  `location`,
  `notes`,
  `created`,
  `updated`
) VALUES (
  NULL,
  'Stephen McAuley',
  'steviebiddles@gmail.com',
  '49.99',
  '2017-12-25 09:00:00',
  '',
  '',
  CURRENT_TIMESTAMP,
  CURRENT_TIMESTAMP
);
