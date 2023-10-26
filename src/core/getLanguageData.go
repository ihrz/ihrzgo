package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func GetLanguage() map[string]interface{} {
	filePath := "src/lang/en-US.yml"

	// Lire le contenu du fichier YAML
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Créez une carte pour stocker les données YAML
	var data map[string]interface{}

	// Analyser le fichier YAML dans la carte
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	return data
}
