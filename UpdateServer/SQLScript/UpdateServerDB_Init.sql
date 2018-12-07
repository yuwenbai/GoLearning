----------------------------------------------------------------------------------------------

----------------------------------------------------------------------------------------------

drop database UpdateServerDB;

create database UpdateServerDB;

use UpdateServerDB;

/*
������
*/

create table UpdateServerDB_Version
(
	ApkName				NVARCHAR (MAX)  NOT  NULL,

	VersionID			NVARCHAR (MAX)  NOT NULL,

	VersionName			NVARCHAR (MAX)   NULL,
		
	VersionInfo			NVARCHAR (MAX)   NULL,
       	
	CreatedOn			DATETIME		 NULL,
)
GO

create table UpdateServerDB_User
(
	UserName				NVARCHAR (MAX)  NOT  NULL,

	PassWord				NVARCHAR (MAX)  NOT NULL,
)
GO

insert into UpdateServerDB_User VALUES("admin","password")
go






