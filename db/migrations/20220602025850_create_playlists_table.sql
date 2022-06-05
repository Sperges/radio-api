-- migrate:up
CREATE TABLE Playlists (
    Id int NOT NULL AUTO_INCREMENT,
    PlaylistName varchar(255) NOT NULL,
    UserId int NOT NULL,
    PRIMARY KEY (Id),
    FOREIGN KEY (UserId) REFERENCES Users(Id)
);

-- migrate:down
DROP TABLE Playlists;
