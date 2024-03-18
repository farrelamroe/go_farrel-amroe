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

#1. Gabungkan data transaksi dari user id 1 dan user id 2
SELECT * FROM Transactions WHERE UserID IN (1, 2);

#2. Tampilkan jumlah harga transaksi user id 1
SELECT SUM(TotalPrice) AS TotalHarga
FROM Transactions
WHERE
    UserID = 1;

#3. Tampilkan total transaksi dengan product type 2
SELECT SUM(TotalPrice) AS TotalTransaksi
FROM
    Transactions
    JOIN TransactionDetails ON Transactions.TransactionID = TransactionDetails.TransactionID
    JOIN Products ON TransactionDetails.ProductID = Products.ProductID
WHERE
    Products.ProductTypeID = 2;

#4. Tampilkan semua field table product dan field name table product type yang saling berhubungan
SELECT Products.*, ProductTypes.Name AS ProductTypeName
FROM Products
    JOIN ProductTypes ON Products.ProductTypeID = ProductTypes.ProductTypeID;

#5. Tampilkan semua field table transaction, field name table product dan field name table user
SELECT Transactions.*, Products.*, Users.*
FROM
    Transactions
    JOIN TransactionDetails ON Transactions.TransactionID = TransactionDetails.TransactionID
    JOIN Products ON TransactionDetails.ProductID = Products.ProductID
    JOIN Users ON Transactions.UserID = Users.UserID;

#6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud
CREATE TRIGGER after_delete_transaction AFTER DELETE 
ON Transactions FOR EACH ROW 
BEGIN 
	DELETE FROM TransactionDetails
	WHERE
	    TransactionID = OLD.TransactionID;
END; 

#7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus
CREATE TRIGGER after_delete_transaction_detail AFTER 
DELETE ON TransactionDetails FOR EACH ROW 
BEGIN 
	UPDATE Transactions
	SET
	    TotalQty = (
	        SELECT SUM(Qty)
	        FROM TransactionDetails
	        WHERE
	            TransactionID = OLD.TransactionID
	    )
	WHERE
	    TransactionID = OLD.TransactionID;
END; 

#8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query
SELECT *
FROM Products
WHERE
    ProductID NOT IN(
        SELECT DISTINCT
            ProductID
        FROM TransactionDetails
    );