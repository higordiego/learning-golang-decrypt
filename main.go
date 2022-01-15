package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func createHash(key string) string {
	// iniciando o modulo de md5
	hasher := md5.New()
	// transformando a string para byte e escrevendo o hash
	hasher.Write([]byte(key))
	// retornando o hash em md5
	return hex.EncodeToString(hasher.Sum(nil))
}

func goRoutines(password, v string, passwordFound chan<- string) {

	// abrindo o arquivo informado
	file, err := os.Open(v)

	// caso não consiga encontrar o arquivo
	if err != nil {
		// informando o erro
		panic(err)
	}

	// lendo o arquivo aberto
	fileScanner := bufio.NewScanner(file)

	// iterando as linhas dos arquivos
	for fileScanner.Scan() {
		// validando o hash é igual a senha passado
		if createHash(fileScanner.Text()) == password {
			// pegando a senha encontrada e passando para o canal
			passwordFound <- fileScanner.Text()
			// parando a iteração
			break
		}
	}
	// fechando o arquivo
	defer file.Close()

}

func main() {

	// a senha para comparação
	password := "1e668de6b2119636f6b37ce07893642d"

	// lista de arquivos
	listFiles := []string{"dictionary1.txt", "dictionary2.txt"}

	// criando um canal para caso a senha seja encontrada
	passwordFound := make(chan string)

	// iterando na lista de arquivos
	for _, v := range listFiles {
		// criando goroutines para comparação de cada lista
		go goRoutines(password, v, passwordFound)
	}

	// informando a senha que está valida.
	fmt.Println("Senha encontrada: ", <-passwordFound)
}
