package toml

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// Method for  read a toml file details
// func ReadTomlFile(pFilename string) (lFileDetails interface{}, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Printf("Panic recovered in ReadTomlFile: %v", r)
// 			err = fmt.Errorf("panic occurred while reading TOML file")
// 		}
// 	}()
// 	_, lErr := toml.DecodeFile(pFilename, &lFileDetails)
// 	if lErr != nil {
// 		log.Println("Error (TRRTF01) ", lErr.Error())
// 		return lFileDetails, lErr
// 	}

// 	return lFileDetails, nil
// }

func ReadTomlFile(filename string) (fileDetails interface{}, err error) {
	log.Println("ReadTomlFile(+)")
	// Decode TOML file
	if _, qerr := toml.DecodeFile(filename, &fileDetails); err != nil {
		log.Printf("Error (TRRTF01) decoding TOML file: %v", qerr)
		return nil, err
	}

	log.Println("ReadTomlFile(-)")
	return fileDetails, nil
}

func GetKeyVal(pData interface{}, key string) string {
	return fmt.Sprintf("%v", pData.(map[string]interface{})[key])
}
