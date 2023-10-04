package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Define os caracteres que podem ser usados na senha
const caracteres = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var EnglishWords = []string{
	"apple",
	"banana",
	"car",
	"dog",
	"elephant",
	"flower",
	"guitar",
	"hat",
	"ice cream",
	"jacket",
	"kite",
	"lion",
	"mountain",
	"notebook",
	"ocean",
	"pencil",
	"quilt",
	"rainbow",
	"star",
	"table",
	"umbrella",
	"violin",
	"waterfall",
	"xylophone",
	"yacht",
	"zebra",
	"airplane",
	"book",
	"computer",
	"desk",
	"elephant",
	"football",
	"garden",
	"hamburger",
	"island",
	"jungle",
	"kangaroo",
	"lemon",
	"monkey",
	"nest",
	"ocean",
	"penguin",
	"quokka",
	"rabbit",
	"sunflower",
	"tiger",
	"umbrella",
	"volcano",
	"watermelon",
	"xylophone",
	"yogurt",
	"zeppelin",
	"ant",
	"butterfly",
	"cat",
	"dolphin",
	"elephant",
	"fox",
	"giraffe",
	"helicopter",
	"iguana",
	"jellyfish",
	"kangaroo",
	"leopard",
	"monkey",
	"narwhal",
	"octopus",
	"panda",
	"quokka",
	"raccoon",
	"seagull",
	"tiger",
	"unicorn",
	"vulture",
	"whale",
	"xylophone",
	"yak",
	"zeppelin",
	"apple",
	"ball",
	"cookie",
	"donkey",
	"eagle",
	"fish",
	"grape",
	"hamburger",
	"ice cream",
	"jacket",
	"koala",
	"lemon",
	"moon",
	"noodle",
	"octopus",
	"penguin",
	"quilt",
	"rainbow",
	"sunflower",
	"tiger",
	"umbrella",
	"violin",
	"watermelon",
	"xylophone",
	"yogurt",
	"zebra",
}

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

func GenerateLogicPassword(leght int) string {
	rand.Seed(time.Now().UnixNano())
	palavraAleatoria := ChoicePassword(leght, EnglishWords)
	return strings.Join(palavraAleatoria, " ")
}

func ChoicePassword(qtdPalavras int, lista []string) []string {
	if qtdPalavras > len(lista) {
		fmt.Println("A quantidade de palavras solicitadas é maior do que a lista disponível.")
		return nil
	}

	// Embaralhe a ordem das palavras na lista original
	rand.Shuffle(len(lista), func(i, j int) {
		lista[i], lista[j] = lista[j], lista[i]
	})

	// Retorne as primeiras qtdPalavras palavras da lista embaralhada
	return lista[:qtdPalavras]
}
