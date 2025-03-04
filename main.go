package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"bufio"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Quantos caracteres você deseja para a senha? ")
	scanner.Scan()
	input := scanner.Text()

	length, err := strconv.Atoi(input)

	if err != nil || length < 8 || length > 32 { 

		fmt.Println("Por favor, insira um número válido entre 8 e 32.")
		return
	}

	password := generatePassword(length)
	fmt.Println("Sua senha gerada é:\n", password)
}

func generatePassword (length int) string{

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"

	allChars:= lowerCase + upperCase + numbers + special

	mandatory :=[]byte{
		lowerCase[rand.Intn(len(lowerCase))],
		upperCase[rand.Intn(len(upperCase))],
		special[rand.Intn(len(special))],
		numbers[rand.Intn(len(numbers))],
	}

	password := make([]byte,length- len(mandatory))

	for i:= range password{
		password[i]=allChars[rand.Intn(len(allChars))]
	}

	password = append(password, mandatory... )

	rand.Shuffle(len(password),func(i, j int) {
		password[i],password[j] = password[j],password[i] 
	})

	return string(password)
}
