DROP TABLE IF EXISTS `UpdateServerDB_Version`;
CREATE TABLE `UpdateServerDB_Version` (
  `ApkName` varchar(60) DEFAULT NULL,
  `VersionID` varchar(60) DEFAULT NULL,
  `VersionName` varchar(60) DEFAULT NULL,
  `VersionInfo` varchar(60) DEFAULT NULL,
  `CreatedOn` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;