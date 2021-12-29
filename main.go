package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func GenerateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ValidateHashMD5(password, text string) {
	if GenerateMD5Hash(text) == password {
		fmt.Println("password:", text)
		os.Exit(0)
	}
	return
}

func goRoutines(password string, phrase chan string) {
	for v := range phrase {
		go ValidateHashMD5(password, v)
	}
}

func main() {

	password := "df6f58808ebfd3e609c234cf2283a989"
	phrase := make(chan string)

	file, err := os.Open("dictionary.txt")

	if err != nil {
		panic(err)
	}

	go goRoutines(password, phrase)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		phrase <- fileScanner.Text()
	}

	close(phrase)

	defer file.Close()

}
