package main

import(
	"database/sql"
	"log"
	"fmt"
	"errors"
)

/**
	Database Operations
**/

var DB_VERSION int

func initDB(){
	DB_VERSION = 0
	db, err := _getDB()
	// Execute the query
	var name string
	err = db.QueryRow("SHOW TABLES LIKE 'Metadata'").Scan(&name)
	switch {
	    case err == sql.ErrNoRows:
	    	DBInstall(db)
	    case err != nil:
	        log.Fatal(err)
	    default:
	        DBVerifyDatabaseVersion(db)
    }
    db.Close()
}

func DBInstall(db  *sql.DB){
	var querystr []string
	switch(DB_VERSION){
	case 0:
		querystr = DB_V_0_1()
	}

	log.Printf("Installing: %s",querystr)
	for _, query := range querystr {
	    err := _DBExcecute(db, query)
		if err != nil{
			log.Fatal(err)
		}
	}
	
}

func DBVerifyDatabaseVersion(db  *sql.DB){
	var Value string
	err := db.QueryRow("select V from Metadata where K = 'DB_VERSION'").Scan(&Value)
	if err != nil{
		log.Fatal(errors.New("Problem with the Instalation, reset Database data sorry "))
	}else{
		if fmt.Sprintf("%d",DB_VERSION) != Value{
			//Do update, when it exists
		}
	}
	
}


func DB_V_0_1() []string {
	return []string {`
		CREATE TABLE ImplementationUser (
		  IdVal int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		  Creation timestamp NULL DEFAULT NULL,
		  Deletion timestamp NULL DEFAULT NULL,
		  UserKey varchar(255) DEFAULT NULL,
		  Status int(11) DEFAULT NULL,
		  ImplementationId varchar(255) DEFAULT NULL
		);`,`
		CREATE TABLE ApiSesion (
		  IdVal int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		  ImplementationUser int(11) DEFAULT NULL,
		  SesionID varchar(255) DEFAULT NULL,
		  Creation timestamp NULL DEFAULT NULL,
		  Used timestamp NULL DEFAULT NULL,
		  Kind int(11) DEFAULT NULL,
		  Status int(11) DEFAULT NULL,
		  Data text
		);`,`
		CREATE TABLE Metadata (
		  IdVal int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
		  K varchar(255) DEFAULT NULL,
		  V varchar(255) DEFAULT NULL,
		  Creation timestamp NULL DEFAULT NULL
		);
		`,"INSERT INTO Metadata VALUES (0,'DB_VERSION','0',now());"}
	}


