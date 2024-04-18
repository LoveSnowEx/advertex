CREATE TABLE IF NOT EXISTS ads
(
    id       INTEGER PRIMARY KEY,
    title    TEXT     NOT NULL,
    start_at DATETIME NOT NULL,
    end_at   DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS ad_age_conditions
(
    id        INTEGER PRIMARY KEY,
    ad_id     INTEGER NOT NULL,
    start_age INTEGER NULL,
    end_age   INTEGER NULL,
    FOREIGN KEY (ad_id) REFERENCES ads (id)
);
CREATE TABLE IF NOT EXISTS ad_gender_conditions
(
    id     INTEGER PRIMARY KEY,
    ad_id  INTEGER NOT NULL,
    gender INTEGER NOT NULL,
    FOREIGN KEY (ad_id) REFERENCES ads (id)
);
CREATE TABLE IF NOT EXISTS ad_country_conditions
(
    id           INTEGER PRIMARY KEY,
    ad_id        INTEGER NOT NULL,
    country_code INTEGER NOT NULL,
    FOREIGN KEY (ad_id) REFERENCES ads (id)
);
CREATE TABLE IF NOT EXISTS ad_platform_conditions
(
    id       INTEGER PRIMARY KEY,
    ad_id    INTEGER NOT NULL,
    platform INTEGER NOT NULL,
    FOREIGN KEY (ad_id) REFERENCES ads (id)
);
