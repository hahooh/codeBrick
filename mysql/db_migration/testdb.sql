USE test_db

CREATE TABLE `inventory` (
    Id INT UNSIGNED AUTO_INCREMENT,
    VehicleIdentificationNumber VARCHAR(50),
    ModelName VARCHAR(50),
    Producer VARCHAR(50),
    Year YEAR(4),
    MSRP DECIMAL(12, 2),
    Status VARCHAR(20),
    Booked BOOLEAN,
    Listed BOOLEAN,
    CreatedDate DATETIME DEFAULT CURRENT_TIMESTAMP,
    ModifiedDate DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`Id`) 
);

INSERT INTO inventory VALUES (1, "MN123402", "320i", "BMW", 2014, 147000.00, "ordered", 1, 1, "2019-09-15 10:45:32", "2019-09-15 10:45:32"),(2, "4JDB1234775", "Camry", "Toyota", 2015, 120000.00, "in stock", 1, 1, "2019-09-15 10:45:32", "2019-09-15 10:45:32"),(3, "TFBA1234274", "Focus", "Ford", 2016, 130000.00, "ordered", 0, 1, "2019-09-15 10:45:32", "2019-09-15 10:45:32"),(4, "G3SBU1234449", "Civic", "Honda", 2016, 140000.00, "sold", 0, 0, "2019-09-15 10:45:32", "2019-09-15 10:45:32");