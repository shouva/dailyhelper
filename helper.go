package dailyhelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(path string) ([]byte, error) {
	filesource, err := os.Open(path)
	if err != nil {
		fmt.Println("convert jsonFile ke byteArray")
		return nil, err
		// panic(err)
	}
	defer filesource.Close()
	return ioutil.ReadAll(filesource)
}

// ReadConfig :
func ReadConfig(path string, config interface{}) error {
	bytefile, err := readFile(path)
	if err != nil {
		fmt.Println("Convert jsonFile ke byteArray gagal")
		return err
		// panic(err)
	}
	err = json.Unmarshal(bytefile, &config)
	if err != nil {
		fmt.Println("Unmarshal config file gagal")
		return err
	}
	return nil
}
