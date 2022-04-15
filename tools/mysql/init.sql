USE `testgo`;

CREATE TABLE `Album`(
    `Id` int not null AUTO_INCREMENT,
    `Title` varchar(150) not null,
    `Artist` varchar(150) not null,
    `Price` decimal(10,2) not null,
    PRIMARY KEY (`Id`)
) AUTO_INCREMENT = 1;

INSERT INTO `testgo`.`Album`  (`Id`, `Title`, `Artist`, `Price`) VALUES ('1', 'Blue train','John Coltrane',56.99);
INSERT INTO `testgo`.`Album`  (`Id`, `Title`, `Artist`, `Price`) VALUES ('2', 'Jeru','Gerry Mulligan',17.99);
INSERT INTO `testgo`.`Album`  (`Id`, `Title`, `Artist`, `Price`) VALUES ('3', 'Sarah Vaughan and Clifford Brown','Sarah Vaughan',39.99);

