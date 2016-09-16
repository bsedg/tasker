CREATE DATABASE IF NOT EXISTS tasker;

CREATE USER IF NOT EXISTS '$MYSQL_USER'@'' IDENTIFIED BY '$MYSQL_USER';

GRANT ALL ON tasker.* TO '$MYSQL_USER'@'';

CREATE TABLE IF NOT EXISTS tasker.tasks (
     id      BIGINT(18) NOT NULL auto_increment,
     name    VARCHAR(255) DEFAULT NULL,
     action  VARCHAR(255) DEFAULT NULL,
     time    VARCHAR(255) DEFAULT NULL,
     created DATETIME DEFAULT NOW(),
     PRIMARY KEY (`id`)
 );
