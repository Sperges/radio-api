-- migrate:up
CREATE TABLE Radios (
    Id int NOT NULL AUTO_INCREMENT,
    RadioName varchar(255) NOT NULL,
    StreamUrl varchar(255) NOT NULL,
    UserId int NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE Radios;
