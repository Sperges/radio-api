-- migrate:up
CREATE TABLE Radios (
    ID int NOT NULL AUTO_INCREMENT,
    RadioName varchar(255) NOT NULL,
    UserID int NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (UserID) REFERENCES Users(ID)
);

-- migrate:down
DROP TABLE Radios;
