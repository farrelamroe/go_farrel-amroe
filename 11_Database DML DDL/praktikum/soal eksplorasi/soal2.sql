CREATE TABLE User (
    UserID INT PRIMARY KEY AUTO_INCREMENT, Username VARCHAR(50) NOT NULL, Email VARCHAR(100) UNIQUE NOT NULL, Password VARCHAR(100) NOT NULL, FullName VARCHAR(100), Address VARCHAR(255), PhoneNumber VARCHAR(20)
);

CREATE TABLE Catatan (
    CatatanID INT PRIMARY KEY AUTO_INCREMENT, UserID INT, Title VARCHAR(255) NOT NULL, Content TEXT, Category VARCHAR(50), FOREIGN KEY (UserID) REFERENCES User (UserID)
);

CREATE TABLE Media (
    MediaID INT PRIMARY KEY AUTO_INCREMENT, CatatanID INT, Type ENUM(
        'image', 'video', 'audio', 'document'
    ), FileURL VARCHAR(255) NOT NULL, FOREIGN KEY (CatatanID) REFERENCES Catatan (CatatanID)
);

CREATE TABLE Task (
    TaskID INT PRIMARY KEY AUTO_INCREMENT, UserID INT, Title VARCHAR(255) NOT NULL, Description TEXT, Status ENUM(
        'pending', 'completed', 'in progress', 'cancelled'
    ), Deadline DATE, FOREIGN KEY (UserID) REFERENCES User (UserID)
);

CREATE TABLE Halaman (
    HalamanID INT PRIMARY KEY AUTO_INCREMENT, UserID INT, Title VARCHAR(255) NOT NULL, Content TEXT, FOREIGN KEY (UserID) REFERENCES User (UserID)
);