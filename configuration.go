package main

import (
	
	"flag"
	"log"
	"fmt"
	"encoding/json"
	"errors"
	"io/ioutil"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

type ConfigStruct struct {
	Inited bool
	Address  string
	Port string
	DBAddress string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	DBTablePrefix string

	Tok3nAPISecret string
	Tok3nAPIKey string
}

var configData ConfigStruct

var configPathFlag = flag.String("path", "", "path to config file")
var configPath string

func configure() error { //Init configuration 
	configData.Inited = false

	configPath = *configPathFlag
	err := _configWithFile(*configPathFlag)
	if err != nil{
		if err.Error() == "Not valid config File." {
			return err;
		}else if err.Error() == "Non existent file." && !configData.Inited{
			configPath = "./opentok3n.config"
			err = _configWithFile(configPath)
			if err!= nil && err.Error() == "Not valid config File." {
				return err;
			}
			_startInteractiveConfiguration()
		}
	}else{
		_startInteractiveConfiguration()
	}
	return nil
}


var portFlag = flag.String("port", "", "server port")
var addrFlag = flag.String("addr", "", "http service address") 
var dbAddrFlag = flag.String("dbaddr", "", "database server address")
var dbPortFlag = flag.String("dbport", "", "database server port")
var dbUserFlag = flag.String("dbuser", "", "database server user")
var dbPassFlag = flag.String("dbpass", "", "database server pass")
var dbDBNameFlag = flag.String("dbname", "", "database name")
var dbTablePrefixFlag = flag.String("dbprefix", "", "database table prefix")

var tok3nSecretFlag = flag.String("tok3nSecret", "", "Tok3n secret API key")
var tok3nApiKeyFlag = flag.String("tok3nKey", "", "Tok3n API key")

func _startInteractiveConfiguration(){
	flag.Parse()
	asked := false

	if configData.Address == ""{ //Not configured the Address
		if *addrFlag == ""{ //Empty flag or not existent
			fmt.Printf("Addres of the current Server: ")
			fmt.Scanf("%s", &configData.Address)
			asked = true
		}else{
			configData.Address = *addrFlag
		}
	}

	if configData.Port == ""{
		if *portFlag == ""{
			fmt.Printf("Port of the current Server: ")
			fmt.Scanf("%s", &configData.Port)
			asked = true
		}else{
			configData.Port = *portFlag
		}
	}

	if configData.DBAddress == ""{
		if *dbAddrFlag == ""{
			fmt.Printf("Address of the MySQL server: ")
			fmt.Scanf("%s", &configData.DBAddress)
			asked = true
		}else{
			configData.DBAddress = *dbAddrFlag
		}
	}

	if configData.DBPort == ""{
		if *dbPortFlag == ""{
			fmt.Printf("Address of the MySQL server (3306): ")
			fmt.Scanf("%s", &configData.DBPort)
			if configData.DBPort == ""{
				configData.DBPort = "3306"
			}
			asked = true
		}else{
			configData.DBPort = *dbPortFlag
		}
	}

	if configData.DBUser == ""{
		if *dbUserFlag == ""{
			fmt.Printf("Database username: ")
			fmt.Scanf("%s", &configData.DBUser)
			asked = true
		}else{
			configData.DBUser = *dbUserFlag
		}
	}

	if configData.DBPassword == ""{
		if *dbPassFlag == ""{
			fmt.Printf("Database password: ")
			fmt.Scanf("%s", &configData.DBPassword)
			asked = true
		}else{
			configData.DBPassword = *dbPassFlag
		}
	}

	if configData.DBName == ""{
		if *dbDBNameFlag == ""{
			fmt.Printf("Database name: ")
			fmt.Scanf("%s", &configData.DBName)
			asked = true
		}else{
			configData.DBName = *dbDBNameFlag
		}
	}

	if configData.DBTablePrefix == ""{
		if *dbTablePrefixFlag == ""{
			fmt.Printf("Database table prefix: ")
			fmt.Scanf("%s", &configData.DBTablePrefix)
			asked = true
		}else{
			configData.DBTablePrefix = *dbTablePrefixFlag
		}
	}

	fmt.Printf("\n\n ---------------------- \n\nGet your FREE API Keys from http://www.tok3n.com/ \n")

	if configData.Tok3nAPISecret == ""{
		if *tok3nSecretFlag == ""{
			fmt.Printf("Tok3n API Secret: ")
			fmt.Scanf("%s", &configData.Tok3nAPISecret)
			asked = true
		}else{
			configData.Tok3nAPISecret = *tok3nSecretFlag
		}
	}

	if configData.Tok3nAPIKey == ""{
		if *tok3nApiKeyFlag == ""{
			fmt.Printf("Tok3n API Key: ")
			fmt.Scanf("%s", &configData.Tok3nAPIKey)
			asked = true
		}else{
			configData.Tok3nAPIKey = *tok3nApiKeyFlag
		}
	}

	configData.Inited = true

	err := verifyMySQLConfiguration()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("\nThe database is correctly loaded.\n")

	if asked{
		var rewrite string
		fmt.Printf("You want to rewrite the configuration file with the previous data (Y/n): ")
		fmt.Scanf("%s", &rewrite)
		if rewrite == "Y"{
			jsonString,_ := json.Marshal(configData)
			fmt.Printf("Ask for save the configuration: %s",jsonString)
			err = ioutil.WriteFile(configPath, jsonString, 0644)
			if err != nil { panic(err) }
		}
	}
	log.Printf("%v",configData)
}

func verifyMySQLConfiguration() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",configData.DBUser, configData.DBPassword,configData.DBAddress,configData.DBPort,configData.DBName))
	if err != nil {
	    return err // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
	    return err // proper error handling instead of panic in your app
	}
	return nil
}

func _configWithFile(filepath string) error{
	configExists := fileExists(filepath)
	if configExists{
		log.Printf("File Exists: %s",filepath)
		err := _parseConfigString(_getContentOfFile(filepath))
		if err != nil{
			return err
		}
	}else{
		return errors.New("Non existent file.")
	}
	return nil
}

func _parseConfigString(config []byte) error {
	err := json.Unmarshal(config, &configData)
	if err != nil {
		return errors.New("Not valid config File.")
	}
	configData.Inited = true

	return nil
}
