-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

CREATE TABLE `ToDo` (
    `ID` bigint(20) NOT NULL AUTO_INCREMENT,
    `Title` varchar(200) DEFAULT NULL,
    `Description` varchar(1024) DEFAULT NULL,
    `Reminder` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`ID`),
    UNIQUE KEY `ID_UNIQUE` (`ID`)
)
