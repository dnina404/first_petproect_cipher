package main

import (
	"XOR/cipherer"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()
	if len(*secretKey) == 0 {
		fmt.Fprintf(os.Stderr, "Ты может ключ введешь?")
		os.Exit(1)
	}
	switch *mode {
	case "cipher":
		plainText := getUserInput("веди че хочешь зашифровать: ")
		cipheredText, err := cipherer.Cipher(plainText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Вот тебе ошибка: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Зашифрованный текст", cipheredText)
	case "decipher":
		cipheredText := getUserInput("Введи че хочешь расшифровать: ")
		cipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Вот тебе ошибка: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Расшифрованный текст", cipheredText)
	default:
		fmt.Println("введи название режима нормально")
		os.Exit(1)
	}
}
func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка какая то вышла молодой человек, введи еще раз")
			continue
		}
		return strings.TrimRight(result, "\n")
	}
}

/*hello
world*/
