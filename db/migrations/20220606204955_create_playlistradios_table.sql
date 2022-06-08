-- migrate:up
CREATE TABLE PlaylistRadios (
    PlaylistId int NOT NULL,
    RadioId int NOT NULL,
    UserId int NOT NULL,
    PRIMARY KEY (PlaylistId, RadioId),
    FOREIGN KEY (PlaylistId) REFERENCES Playlists(Id) ON DELETE CASCADE,
    FOREIGN KEY (RadioId) REFERENCES Radios(Id) ON DELETE CASCADE,
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE Radios;
