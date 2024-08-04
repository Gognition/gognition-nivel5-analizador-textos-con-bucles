package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	texto := `Go es un lenguaje de programación concurrente y compilado 
    inspirado en la sintaxis de C. Ha sido desarrollado por Google, y sus 
    diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson. 
    Go es eficiente, escalable y productivo. También maneja texto en español: á, é, í, ó, ú.`

	fmt.Println("Analizador de Texto Multifuncional")
	fmt.Println("----------------------------------")

	fmt.Printf("1. Número total de palabras: %d\n", contarPalabras(texto))
	fmt.Printf("2. Palabra más larga: %s\n", palabraMasLarga(texto))
	mostrarFrecuenciaLetras(texto)
	fmt.Printf("4. Texto con palabras invertidas: %s\n", invertirPalabras(texto))
	fmt.Printf("5. Número de palabras clave: %d\n", buscarPalabrasClave(texto, []string{"Go", "programación", "Google"}))
	fmt.Println("6. Análisis de estructura de oraciones:")
	analizarEstructuraOraciones(texto)
}

func contarPalabras(texto string) int {
	palabras := strings.Fields(texto)
	count := 0
	for i := 0; i < len(palabras); i++ {
		count++
	}
	return count
}

func palabraMasLarga(texto string) string {
	palabras := strings.Fields(texto)
	masLarga := ""
	for _, palabra := range palabras {
		if len(palabra) > len(masLarga) {
			masLarga = palabra
		}
	}
	return masLarga
}

func mostrarFrecuenciaLetras(texto string) {
	frecuencia := make(map[rune]int)
	textoRunes := []rune(texto)
	i := 0

	fmt.Println("3. Frecuencia de letras:")

	// Ejemplo de cómo 'len' puede cambiar con caracteres especiales
	palabraNormal := "cafeteria"
	palabraConTilde := "cafetería"
	fmt.Printf("   Longitud de '%s': %d\n", palabraNormal, len(palabraNormal))
	fmt.Printf("   Longitud de '%s': %d\n", palabraConTilde, len(palabraConTilde))
	fmt.Printf("   Longitud real de '%s': %d\n", palabraConTilde, len([]rune(palabraConTilde)))

	// Bucle principal para contar frecuencia de letras
	for i < len(textoRunes) {
		letra := textoRunes[i]
		letraLower := unicode.ToLower(letra)
		if unicode.IsLetter(letraLower) {
			frecuencia[letraLower]++
		}
		i++
	}

	// Mostrar resultados
	for letra, freq := range frecuencia {
		fmt.Printf("   %c: %d\n", letra, freq)
	}

	fmt.Printf("   Total de letras únicas: %d\n", len(frecuencia))
}

func invertirPalabras(texto string) string {
	palabras := strings.Fields(texto)
	for i, palabra := range palabras {
		runes := []rune(palabra)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		palabras[i] = string(runes)
	}
	return strings.Join(palabras, " ")
}

func buscarPalabrasClave(texto string, claves []string) int {
	palabras := strings.Fields(strings.ToLower(texto))
	count := 0
	for _, palabra := range palabras {
		for _, clave := range claves {
			if palabra == strings.ToLower(clave) {
				count++
				break // Evita contar la misma palabra más de una vez
			}
		}
	}
	return count
}

func analizarEstructuraOraciones(texto string) {
	oraciones := strings.Split(texto, ".")
	for i := 0; ; i++ {
		if i >= len(oraciones) {
			break
		}
		oracion := strings.TrimSpace(oraciones[i])
		if oracion == "" {
			continue
		}
		palabras := strings.Fields(oracion)
		fmt.Printf("   Oración %d: %d palabras\n", i+1, len(palabras))
	}
}
