CREATE TABLE `lists` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `organiser` varchar(128) NOT NULL DEFAULT '',
  `email` varchar(128) NOT NULL DEFAULT '',
  `amount` double(8,2) NOT NULL,
  `date` datetime NOT NULL,
  `location` varchar(255) DEFAULT '',
  `notes` varchar(255) DEFAULT '',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;