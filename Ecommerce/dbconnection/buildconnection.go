package dbconnection

import "log"

func BuildConnection() error {
	log.Println("BuildConnection (+) ")
	var lErr error

	Gdb_instance.Gormdb, Gdb_instance.Mysql_sqldb, lErr = Dbconnection()
	if lErr != nil {
		log.Println("Error (DCBC01) ", lErr.Error())
		return lErr
	}
	log.Println("BuildConnection (-) ")
	return lErr
}
