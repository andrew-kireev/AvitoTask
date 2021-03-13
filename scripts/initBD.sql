CREATE TABLE IF NOT EXISTS Statistic
(
    stat_date Date NOT NULL,
    views     INTEGER,
    clicks    INTEGER,
    cost      decimal,
    cpc       decimal,
    cpm       decimal
);