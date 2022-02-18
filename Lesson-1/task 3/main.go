package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getEnvString(env string, defaultVault string) string {
	if envVal, ok := os.LookupEnv(env); ok {
		return envVal
	}
	return defaultVault
}

func getEnvInt64(env string, defaultVault int64) int64 {
	if envVal, ok := os.LookupEnv(env); ok {
		envInt64, err := strconv.ParseInt(envVal, 10, 64)
		if err == nil {
			return envInt64
		}
	}
	return defaultVault
}

var (
	ErrNumberFileNegative = errors.New("number of files cannot be negative")
	ErrFileCreate         = errors.New("file creation failed")
	ErrFileClose          = errors.New("file close error")
	ErrWriteFile          = errors.New("file write error")
	ErrReadDirectory      = errors.New("directory read error")
	ErrCreateDirectory    = errors.New("directory create error")
)

func main() {
	pathDir := flag.String("dir", getEnvString("PATH_DIRECTORY", "/tmp/lesson-1"), "path to the directory, where the files will be created")
	numberFiles := flag.Int64("n", getEnvInt64("NUMBER_FILES", 1), "number of files created")
	flag.Parse()

	_, err := ioutil.ReadDir(*pathDir)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrReadDirectory, err.Error())
		fmt.Println(err)
	}

	if errors.Is(err, ErrReadDirectory) {
		err = os.Mkdir(*pathDir, 0777)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrCreateDirectory, err.Error())
			fmt.Println(err)
		}
	}

	if *numberFiles < 0 {
		fmt.Println(ErrNumberFileNegative)
		return
	}

	var fd *os.File
	defer func() {
		err = fd.Close()
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrFileClose, err.Error())
			fmt.Println(err)
			return
		}
	}()

	for i := int64(1); i <= *numberFiles; i++ {
		fd, err = os.Create(fmt.Sprintf("%s/file-%d.txt", *pathDir, i))
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrFileCreate, err.Error())
			fmt.Println(err)
			break
		}

		_, err = fmt.Fprintln(fd, fmt.Sprintf("data: %s/file-%d.txt", *pathDir, i))
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrWriteFile, err)
			fmt.Println(err)
			break
		}
		fmt.Printf("file-%d.txt created\n", i)
	}
}
