// showmetar.go
package main

import (
	"fmt"
	"os"

	"github.com/danielribes/metar"
)

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	metarline, data, descripcio := metar.GetMETAR(os.Args[1])

	fmt.Printf("%s\n", descripcio)
	fmt.Printf("%s\n", metarline)
	fmt.Printf("Ultima actualització: %s\n", data)
}

func help() {
	fmt.Printf("metar 0.1\n")
	fmt.Printf("--> Cal indicar un codi OACI d'aeroport\n\n")
	fmt.Printf("Exemple: metar LELL\n")
}
