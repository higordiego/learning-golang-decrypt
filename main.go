package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func GenerateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GenerateSHA1Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateHashShA1(password, text string) {
	if GenerateSHA1Hash(text) == password {
		fmt.Println("Password crack finish...")
		fmt.Println("cryptography: SHA1\npassword: ", text)
		os.Exit(0)
	}
	return

}

func ValidateHashMD5(password, text string) {
	if GenerateMD5Hash(text) == password {
		fmt.Println("Password crack finish..")
		fmt.Println("cryptography: MD5\npassword:", text)
		os.Exit(0)
	}
	return
}

func goRoutines(password string, phrase chan string) {
	for v := range phrase {
		go ValidateHashMD5(password, v)
		go ValidateHashShA1(password, v)
	}
}

func main() {
	password := flag.String("h", "", "hash em md5. (Required)")
	flag.Parse()

	if *password == "" {
		return
	}

	phrase := make(chan string)

	file, err := os.Open("dictionary.txt")

	if err != nil {
		panic(err)
	}

	go goRoutines(*password, phrase)

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		phrase <- fileScanner.Text()
	}

	close(phrase)

	defer file.Close()

}
