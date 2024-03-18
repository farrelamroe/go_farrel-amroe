CREATE TABLE Operators (
    OperatorID INT(11) PRIMARY KEY AUTO_INCREMENT, Name VARCHAR(255), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP
);

CREATE TABLE ProductTypes (
    ProductTypeID INT(11) PRIMARY KEY AUTO_INCREMENT, Nama VARCHAR(255), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP
);

CREATE TABLE ProductDescriptions (
    ProductDescriptionID INT(11) PRIMARY KEY AUTO_INCREMENT, Description TEXT, CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP
);

CREATE TABLE Products (
    ProductID INT(11) PRIMARY KEY AUTO_INCREMENT, Code VARCHAR(50), Name VARCHAR(100), Status SMALLINT, ProductTypeID INT(11), OperatorID INT(11), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP, FOREIGN KEY (ProductTypeID) REFERENCES ProductTypes (ProductTypeID), FOREIGN KEY (OperatorID) REFERENCES Operators (OperatorID)
);

CREATE TABLE PaymentMethods (
    PaymentMethodID INT(11) PRIMARY KEY AUTO_INCREMENT, Name VARCHAR(255), Status SMALLINT, CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP
);

CREATE TABLE Users (
    UserID INT(11) PRIMARY KEY AUTO_INCREMENT, Status SMALLINT, Dob DATE, Gender CHAR(1), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP
);

CREATE TABLE TransactionDetails (
    TransactionDetailID INT(11) PRIMARY KEY AUTO_INCREMENT, TransactionID INT(11), ProductID INT(11), Status VARCHAR(10), Qty INT(11), Price NUMERIC(25, 2), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP, FOREIGN KEY (ProductID) REFERENCES Products (ProductID)
);

CREATE TABLE Transactions (
    TransactionID INT(11) PRIMARY KEY AUTO_INCREMENT, UserID INT(11), PaymentMethodID INT(11), Status VARCHAR(10), TotalPrice NUMERIC(25, 2), CreatedAt TIMESTAMP, UpdatedAt TIMESTAMP, FOREIGN KEY (UserID) REFERENCES Users (UserID), FOREIGN KEY (PaymentMethodID) REFERENCES PaymentMethods (PaymentMethodID)
);

#a. Delete data pada tabel product dengan id = 1
DELETE FROM Products WHERE ProductID = 1;

#b. Delete data pada tabel product dengan product type id 1
DELETE FROM Products WHERE ProductTypeID = 1;