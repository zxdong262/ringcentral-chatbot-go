package ringcentralchatbot

import (
	"log"
)

const folderName string = "fieldb"

func DbAction(tableName string, action string, data string) {
	log.Println(tableName, action, data)
	return nil
}
