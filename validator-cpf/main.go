package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func removeCharacters(cleanCPF string) string {
	re := regexp.MustCompile(`[^\d]+`)
	cpf := re.ReplaceAllString(cleanCPF, "")
	return cpf
}

func validateCPF(c string) bool {
	//remove caracteres como .(ponto) e -(traço)
	cpf := removeCharacters(c)

	//verifica se o CPF digitado tem 11 digitos, pois ja removi . -
	if len(cpf) != 11 {
		return false
	}

	//verifica se todos os dígitos são iguais
	if cpf[0] == cpf[1] && cpf[1] == cpf[2] && cpf[2] == cpf[3] && cpf[3] == cpf[4] && cpf[4] == cpf[5] && cpf[5] == cpf[6] && cpf[6] == cpf[7] && cpf[7] == cpf[8] && cpf[8] == cpf[9] && cpf[9] == cpf[10] && cpf[10] == cpf[0] {
		return false
	}

	//calculando primeiro dígito verificador
	sum1 := 0
	for i := range 9 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum1 += digit * (10 - i)
	}

	rest1 := (sum1 * 10) % 11
	if rest1 == 10 {
		rest1 = 0
	}

	firstDigit, err := strconv.Atoi(string(cpf[9]))
	if err != nil {
		return false
	}

	if rest1 != firstDigit {
		return false
	}

	//calculando o segundo dígito verificador
	sum2 := 0
	for i := range 10 {
		digit, err := strconv.Atoi(string(cpf[i]))
		if err != nil {
			return false
		}
		sum2 += digit * (11 - i)
	}

	rest2 := (sum2 * 10) % 11
	if rest2 == 10 {
		rest2 = 0
	}

	secondDigit, err := strconv.Atoi(string(cpf[10]))
	if err != nil {
		return false
	}

	if rest2 != secondDigit {
		return false
	}

	// cpf válido
	return true
}

func main() {
	var cpf string
	fmt.Println("Digite o CPF:")
	fmt.Scan(&cpf)

	if validateCPF(cpf) {
		fmt.Println("CPF VÁLIDO !!")
	} else {
		fmt.Println("CPF INVÁLIDO !!")
	}
}

//*                      ________  ________  ________ 
//*                     |\   ____\|\   __  \|\  _____\
//*                     \ \  \___|\ \  \|\  \ \  \__/ 
//*                      \ \  \    \ \   ____\ \   __\
//*                       \ \  \____\ \  \___|\ \  \_|
//*                        \ \_______\ \__\    \ \__\ 
//*                         \|_______|\|__|     \|__| 
//*                                                  
                               
                              