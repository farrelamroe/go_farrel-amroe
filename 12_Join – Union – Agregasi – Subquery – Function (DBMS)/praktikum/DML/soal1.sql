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

# a. Insert 5 operators pada tabel Operators:
INSERT INTO
    Operators (Name, CreatedAt, UpdatedAt)
VALUES ('Operator 1', NOW(), NOW()),
    ('Operator 2', NOW(), NOW()),
    ('Operator 3', NOW(), NOW()),
    ('Operator 4', NOW(), NOW()),
    ('Operator 5', NOW(), NOW());

#b. Insert 3 product type pada tabel ProductTypes:
INSERT INTO
    ProductTypes (Nama, CreatedAt, UpdatedAt)
VALUES ('Type 1', NOW(), NOW()),
    ('Type 2', NOW(), NOW()),
    ('Type 3', NOW(), NOW());

#c. Insert 2 product dengan product type id = 1, dan operator id = 3:
INSERT INTO
    Products (
        Code, Name, Status, ProductTypeID, OperatorID, CreatedAt, UpdatedAt
    )
VALUES (
        'PROD001', 'Product 1', 1, 1, 3, NOW(), NOW()
    ),
    (
        'PROD002', 'Product 2', 1, 1, 3, NOW(), NOW()
    );

#d. Insert 3 product dengan product type id = 2, dan operator id = 1:
INSERT INTO
    Products (
        Code, Name, Status, ProductTypeID, OperatorID, CreatedAt, UpdatedAt
    )
VALUES (
        'PROD003', 'Product 3', 1, 2, 1, NOW(), NOW()
    ),
    (
        'PROD004', 'Product 4', 1, 2, 1, NOW(), NOW()
    ),
    (
        'PROD005', 'Product 5', 1, 2, 1, NOW(), NOW()
    );

#e. Insert 3 product dengan product type id = 3, dan operator id = 4:
INSERT INTO
    Products (
        Code, Name, Status, ProductTypeID, OperatorID, CreatedAt, UpdatedAt
    )
VALUES (
        'PROD006', 'Product 6', 1, 3, 4, NOW(), NOW()
    ),
    (
        'PROD007', 'Product 7', 1, 3, 4, NOW(), NOW()
    ),
    (
        'PROD008', 'Product 8', 1, 3, 4, NOW(), NOW()
    );

#f. Insert product description pada setiap product:
INSERT INTO
    ProductDescriptions (
        Description, CreatedAt, UpdatedAt
    )
VALUES (
        'Description for Product 1', NOW(), NOW()
    ),
    (
        'Description for Product 2', NOW(), NOW()
    ),
    (
        'Description for Product 3', NOW(), NOW()
    ),
    (
        'Description for Product 4', NOW(), NOW()
    ),
    (
        'Description for Product 5', NOW(), NOW()
    ),
    (
        'Description for Product 6', NOW(), NOW()
    ),
    (
        'Description for Product 7', NOW(), NOW()
    ),
    (
        'Description for Product 8', NOW(), NOW()
    );

#g. Insert 3 payment methods
INSERT INTO
    PaymentMethods (
        Name, Status, CreatedAt, UpdatedAt
    )
VALUES ('Method 1', 1, NOW(), NOW()),
    ('Method 2', 1, NOW(), NOW()),
    ('Method 3', 1, NOW(), NOW());

#h. Insert 5 user pada tabel Users
INSERT INTO
    Users (
        Status, Dob, Gender, CreatedAt, UpdatedAt
    )
VALUES (
        1, '2000-01-01', 'M', NOW(), NOW()
    ),
    (
        1, '2000-02-01', 'F', NOW(), NOW()
    ),
    (
        1, '2000-03-01', 'M', NOW(), NOW()
    ),
    (
        1, '2000-04-01', 'F', NOW(), NOW()
    ),
    (
        1, '2000-05-01', 'M', NOW(), NOW()
    );

#i. Insert 3 transaksi di masing-masing user
-- User 1
INSERT INTO
    Transactions (
        UserID, PaymentMethodID, Status, TotalPrice, CreatedAt, UpdatedAt
    )
VALUES (
        1, 1, 'Success', 100.00, NOW(), NOW()
    ),
    (
        1, 2, 'Pending', 150.00, NOW(), NOW()
    ),
    (
        1, 3, 'Failed', 200.00, NOW(), NOW()
    );

-- User 2
INSERT INTO
    Transactions (
        UserID, PaymentMethodID, Status, TotalPrice, CreatedAt, UpdatedAt
    )
VALUES (
        2, 1, 'Success', 110.00, NOW(), NOW()
    ),
    (
        2, 2, 'Success', 120.00, NOW(), NOW()
    ),
    (
        2, 3, 'Pending', 130.00, NOW(), NOW()
    );

-- User 3
INSERT INTO
    Transactions (
        UserID, PaymentMethodID, Status, TotalPrice, CreatedAt, UpdatedAt
    )
VALUES (
        3, 1, 'Pending', 120.00, NOW(), NOW()
    ),
    (
        3, 2, 'Success', 130.00, NOW(), NOW()
    ),
    (
        3, 3, 'Success', 140.00, NOW(), NOW()
    );

-- User 4
INSERT INTO
    Transactions (
        UserID, PaymentMethodID, Status, TotalPrice, CreatedAt, UpdatedAt
    )
VALUES (
        4, 1, 'Success', 150.00, NOW(), NOW()
    ),
    (
        4, 2, 'Failed', 160.00, NOW(), NOW()
    ),
    (
        4, 3, 'Success', 170.00, NOW(), NOW()
    );

-- User 5
INSERT INTO
    Transactions (
        UserID, PaymentMethodID, Status, TotalPrice, CreatedAt, UpdatedAt
    )
VALUES (
        5, 1, 'Success', 180.00, NOW(), NOW()
    ),
    (
        5, 2, 'Success', 190.00, NOW(), NOW()
    ),
    (
        5, 3, 'Success', 200.00, NOW(), NOW()
    );

# j. Insert 3 product di masing-masing transaksi
-- Transaksi 1
INSERT INTO
    TransactionDetails (
        TransactionID, ProductID, Status, Qty, Price, CreatedAt, UpdatedAt
    )
VALUES (
        1, 1, 'Success', 2, 50.00, NOW(), NOW()
    ),
    (
        1, 2, 'Success', 3, 100.00, NOW(), NOW()
    );

-- Transaksi 2
INSERT INTO
    TransactionDetails (
        TransactionID, ProductID, Status, Qty, Price, CreatedAt, UpdatedAt
    )
VALUES (
        2, 3, 'Pending', 1, 110.00, NOW(), NOW()
    ),
    (
        2, 4, 'Success', 2, 120.00, NOW(), NOW()
    );

-- Transaksi 3
INSERT INTO
    TransactionDetails (
        TransactionID, ProductID, Status, Qty, Price, CreatedAt, UpdatedAt
    )
VALUES (
        3, 5, 'Failed', 1, 120.00, NOW(), NOW()
    ),
    (
        3, 6, 'Success', 2, 130.00, NOW(), NOW()
    );