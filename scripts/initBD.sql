create database statistic_bd;
CREATE USER andrew WITH ENCRYPTED PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE statistic_bd TO andrew;

CREATE TABLE IF NOT EXISTS Statistic
(
    stat_date Date NOT NULL,
    views     INTEGER,
    clicks    INTEGER,
    cost      decimal,
    cpc       decimal,
    cpm       decimal
);