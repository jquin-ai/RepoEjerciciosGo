package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"tp0/ejercicios"
)

func leerArchivoComoSlice(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("no se pudo abrir %s: %v", path, err)
	}
	defer file.Close()

	var numeros []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		n, err := strconv.Atoi(linea)
		if err != nil {
			log.Fatalf("línea inválida en %s: %q (%v)", path, linea, err)
		}
		numeros = append(numeros, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error leyendo %s: %v", path, err)
	}
	return numeros
}

func main() {
	a := leerArchivoComoSlice("archivo1.in")
	b := leerArchivoComoSlice("archivo2.in")

	mayor := a
	if ejercicios.Comparar(a, b) == -1 {
		mayor = b
	}

	ejercicios.Seleccion(mayor)

	for _, n := range mayor {
		fmt.Println(n)
	}
}
