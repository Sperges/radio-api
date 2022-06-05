-- migrate:up
CREATE TABLE Users (
    Id int NOT NULL AUTO_INCREMENT,
    Username varchar(255) NOT NULL,
    PRIMARY KEY (Id)
);

-- migrate:down
DROP TABLE Users;
