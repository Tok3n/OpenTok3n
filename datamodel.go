package main
import (
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
	err = _DBExcecute(db,fmt.Sprintf("INSERT INTO ApiSesion VALUES (0,0,'%s',now(),NULL,3, -1, '%s')",session,data))
	db.Close()
	return err
}