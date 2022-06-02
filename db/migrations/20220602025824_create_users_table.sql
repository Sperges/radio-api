-- migrate:up
CREATE TABLE Users (
    ID int NOT NULL AUTO_INCREMENT,
    UserName varchar(255) NOT NULL,
    PRIMARY KEY (ID)
);

-- migrate:down
DROP TABLE Users;
