package main
import (
	"log"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

/**
Session
**/

type ApiSesion struct{
	IdVal	int
    ImplementationUser int
	SesionID	string
	Creation	time.Time
	Used	time.Time
    Kind int // 1:auth; 2:access; 3:Tok3naccess
    Status int //auth{0:issued,1:valid,2:invalid}
    Data string
}


func addTok3nAuthenticationSeccion(session string,data string)error{
	db,err := _getDB()
	if err != nil{
		return err
	}
	err = _DBExcecute(db,fmt.Sprintf("INSERT INTO ApiSesion VALUES (0,0,'%s',now(),DATE_ADD(NOW(), INTERVAL 100 YEAR),3, 0, '%s')",session,data))
	db.Close()
	return err
}

func setValidSession(key string)error{
	db,err := _getDB()
	if err != nil{
		return err
	}
	err = _DBExcecute(db,fmt.Sprintf("UPDATE ApiSesion SET Status = 1 WHERE UserKey = '%s'",key))
	db.Close()
	return err
}
func getApiSessionWithSesionId(session string) (ApiSesion, error) {
	db,err := _getDB()
	var sess ApiSesion
	query := fmt.Sprintf("SELECT * FROM ApiSesion WHERE SesionID = '%s'",session)
	err = db.QueryRow(query).Scan(&sess.IdVal,&sess.ImplementationUser,&sess.SesionID,&sess.Creation,&sess.Used,&sess.Kind,&sess.Status,&sess.Data)
	return sess, err
}

/**
UserImplementation
**/

type ImplementationUser struct {
	IdVal int
	Creation time.Time
    Deletion time.Time
	UserKey string
	Status int //-1: Deleted; 2:Acepted
	Data string
}

func userKeyExists(key string)bool{
	db,err := _getDB()
	if err != nil{
		return false
	}
	var count int
	err = db.QueryRow(fmt.Sprintf("SELECT count(*) as d FROM ImplementationUser WHERE UserKey = '%s'",key)).Scan(&count)
	if err != nil{
		log.Printf(err.Error())
		return false
	}
	return count != 0
}

func createImplementationUser(key,data string) error{
	db,err := _getDB()
	if err != nil{
		return err
	}
	err = _DBExcecute(db,fmt.Sprintf("INSERT INTO ImplementationUser VALUES (0,now(),DATE_ADD(NOW(), INTERVAL 100 YEAR), '%s', 2, '%s')",key,data))
	db.Close()
	return err
}