CREATE TABLE IF NOT EXISTS Statistic
(
    stat_date Date NOT NULL,
    views     INTEGER,
    clicks    INTEGER,
    cost      money,
    cpc       money,
    cpm       money
);