package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Faturamento struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {
	print("N1\n")
	n1()
	print("\n\n")

	print("N2\n")
	n2()
	print("\n\n")

	print("N3\n")
	n3()
	print("\n\n")

	print("N4\n")
	n4()
	print("\n\n")

	print("N5\n")
	n5()
}

func n1() {
	INDICE := 13
	SOMA := 0
	K := 0

	for K < INDICE {
		K += 1
		SOMA += K
	}

	fmt.Printf("A soma dos %d primeiros números é %d.\n", INDICE, SOMA)
}

func n2() {
	numeros := []int{8, 10, 21}

	for _, num := range numeros {
		if pertenceFibonacci(num) {
			fmt.Printf("%d pertence a sequência de Fibonacci.\n", num)
		} else {
			fmt.Printf("%d não pertence a sequência de Fibonacci.\n", num)
		}
	}
}

func pertenceFibonacci(n int) bool {
	a, b := 0, 1
	for a <= n {
		if a == n {
			return true
		}
		a, b = b, a+b
	}
	return false
}

func n3() {
	faturamentos := carregarDados("dados.json")

	menor, maior := calcularMenorMaiorFaturamento(faturamentos)
	diasAcimaMedia := calcularDiasAcimaMedia(faturamentos)

	fmt.Printf("Menor faturamento diário: %.2f\n", menor)
	fmt.Printf("Maior faturamento diário: %.2f\n", maior)
	fmt.Printf("Número de dias com faturamento acima da média: %d\n", diasAcimaMedia)
}

func carregarDados(caminho string) []Faturamento {
	arquivo, err := os.Open(caminho)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	bytes, err := io.ReadAll(arquivo)
	if err != nil {
		log.Fatal(err)
	}

	var faturamentos []Faturamento
	if err := json.Unmarshal(bytes, &faturamentos); err != nil {
		log.Fatal(err)
	}

	return faturamentos
}

func calcularMenorMaiorFaturamento(faturamentos []Faturamento) (float64, float64) {
	var menor, maior float64
	inicializado := false

	for _, fat := range faturamentos {
		if fat.Valor > 0 {
			if !inicializado {
				menor, maior = fat.Valor, fat.Valor
				inicializado = true
			} else {
				if fat.Valor < menor {
					menor = fat.Valor
				}
				if fat.Valor > maior {
					maior = fat.Valor
				}
			}
		}
	}

	return menor, maior
}

func calcularDiasAcimaMedia(faturamentos []Faturamento) int {
	var soma float64
	var diasComFaturamento int

	for _, fat := range faturamentos {
		if fat.Valor > 0 {
			soma += fat.Valor
			diasComFaturamento++
		}
	}

	if diasComFaturamento == 0 {
		return 0
	}

	media := soma / float64(diasComFaturamento)
	diasAcimaMedia := 0

	for _, fat := range faturamentos {
		if fat.Valor > media {
			diasAcimaMedia++
		}
	}

	return diasAcimaMedia
}

func n4() {
	faturamentos := []float64{67836.43, 36678.66, 29229.88, 27165.48, 19849.53}
	estados := []string{"SP", "RJ", "MG", "ES", "Outros"}

	total := calcularTotal(faturamentos)

	fmt.Println("Percentual de representação por estado:")
	for i, valor := range faturamentos {
		percentual := (valor / total) * 100
		fmt.Printf("%s: %.2f%%\n", estados[i], percentual)
	}
}

func calcularTotal(valores []float64) float64 {
	total := 0.0
	for _, v := range valores {
		total += v
	}
	return total
}

func n5() {
	str := "Target Sistemas"
	invertida := inverterString(str)

	fmt.Printf("Original: %s\nInvertida: %s\n", str, invertida)
}

func inverterString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
