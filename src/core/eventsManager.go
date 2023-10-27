package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

func LoadHandlersFromDirectory(dir string) {
	// Lister les fichiers dans le répertoire
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du répertoire:", err)
		return
	}

	// Parcourir les fichiers
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			// Construire le chemin complet du fichier
			filePath := filepath.Join(dir, file.Name())

			// Charger dynamiquement le package et les fonctions
			loadHandlersFromFile(filePath)
		}
	}
}

func loadHandlersFromFile(filePath string) {
	// Charger dynamiquement le fichier Go
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier %s : %s\n", filePath, err)
		return
	}
	defer file.Close()

	// Vous pouvez utiliser un outil de parsing Go comme "go/parser" ou "go/ast" pour extraire
	// les fonctions ou objets du fichier, mais cela peut être complexe.
	// Une autre option consiste à charger le code comme une bibliothèque et
	// à utiliser la réflexion pour rechercher les fonctions dans le package events.
	packagePath := "events" // Nom du package dans les fichiers

	// Charger le package dynamiquement
	pkg, err := plugin.Open(filePath)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du package %s : %s\n", packagePath, err)
		return
	}

	// Utiliser la réflexion pour obtenir les fonctions du package
	sym, err := pkg.Lookup("FunctionName") // Remplacez "FunctionName" par le nom de la fonction à charger
	if err != nil {
		fmt.Printf("Erreur lors de la recherche de la fonction : %s\n", err)
		return
	}

	// Assurez-vous que le symbole est une fonction
	if f, ok := sym.(func()); ok {
		// Appeler la fonction chargée dynamiquement
		f()
	} else {
		fmt.Printf("Le symbole n'est pas une fonction\n")
	}
}
