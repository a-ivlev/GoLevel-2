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
	ErrNubberFileNegative = errors.New("number of files cannot be negative")
	ErrFileCreate         = errors.New("file creation failed")
	ErrFileClose          = errors.New("file close error")
	ErrWriteFile          = errors.New("file write error")
	ErrReadDirectory      = errors.New("directory read error")
	ErrCreateDirectory    = errors.New("directory create error")
)

func main() {

}
