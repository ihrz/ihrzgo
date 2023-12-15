package lang

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func GetLanguage(guildId *string) map[string]interface{} {
	// faire une requests a la db pour savoir en quelle langue le server la configurée

	filePath := "src/lang/en-US.yml"

	// var defaultLanguage = "en-US"
	// var lang = db.GetLanguagePerGuildID(guildId)

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
