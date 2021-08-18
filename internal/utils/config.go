package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

const ConfigFileName = "../../configs/config.yaml"

func LoadConfigTenTimes() {
	for i := 0; i < 10; i++ {
		LoadConfig(ConfigFileName)
	}
}

func LoadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panicf("can't open file %s: %v", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Can't close file %s\n", fileName)
		}
	}()

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Panicf("can't read file contents %s: %v", fileName, err)
	}

	fmt.Printf("File contents are: \n%s\n", string(contents))
}
