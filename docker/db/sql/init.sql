CREATE USER 'migration-user'@'%' IDENTIFIED WITH 'mysql_native_password' BY 'local-dev' REQUIRE NONE PASSWORD EXPIRE DEFAULT ACCOUNT UNLOCK;
GRANT ALTER, CREATE, CREATE VIEW, DROP, INDEX, REFERENCES ON `local`.* TO 'migration-user'@'%';
GRANT SELECT, INSERT, DELETE, UPDATE ON `local`.* TO 'migration-user'@'%';