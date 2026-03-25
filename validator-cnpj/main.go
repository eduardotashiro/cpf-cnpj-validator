package main

import (
	"fmt"
	"regexp"
	// "strconv"
)

func removeCharacters(cleanCNPJ string) string {
	re := regexp.MustCompile(`[^\d]+`)
	cnpj := re.ReplaceAllString(cleanCNPJ, "")
	return cnpj
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func validateCnpj(c string) bool {
	cnpj := removeCharacters(c)

	if len(cnpj) != 14 {
		return false
	}

	counter := 0
	for k := 0; k < len(cnpj); k++ {
		if cnpj[0] == cnpj[k] {
			counter++
		}
		if counter == len(cnpj) {
			return false
		}
	}

	// calcular aqui

	return true
}

func main() {

	var CNPJ string = "59.541.264/0001-03"
	fmt.Println(CNPJ)
	cnpjInv := reverse(CNPJ)
	fmt.Println(cnpjInv)

}
