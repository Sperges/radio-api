-- migrate:up
CREATE TABLE Playlists (
    ID int NOT NULL AUTO_INCREMENT,
    PlaylistName varchar(255) NOT NULL,
    UserID int NOT NULL,
    PRIMARY KEY (ID),
    FOREIGN KEY (UserID) REFERENCES Users(ID)
);

-- migrate:down
DROP TABLE Playlists;
