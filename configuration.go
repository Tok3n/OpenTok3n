package main

import (
	
	"flag"
	"log"
	//"fmt"
	"encoding/json"
	"errors"
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
			if err.Error() == "Not valid config File." {
				return err;
			}
			_startInteractiveConfiguration()
		}
	}else{
		_startInteractiveConfiguration()
	}
	return nil
}

func _startInteractiveConfiguration(){
	log.Printf("%v",configData)
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
