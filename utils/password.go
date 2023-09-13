package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Define os caracteres que podem ser usados na senha
const caracteres = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// Função para gerar uma senha aleatória
func GeneratePassword(leght int) string {
	if leght < 8 {
		fmt.Println("Your Password is not secure")
		os.Exit(1)
	}
	rand.Seed(time.Now().UnixNano())
	senha := make([]byte, leght)
	for i := 0; i < leght; i++ {
		senha[i] = caracteres[rand.Intn(len(caracteres))]
	}
	return string(senha)
}
