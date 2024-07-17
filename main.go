package main

import (
	"os"
	"text/template"
	"log"
)

type ServerProperties struct {
	MOTD             string
	Memory           string
	Port             string
	Seed             string
	GameMode         string
	Difficulty       string
	WhiteList        string
	WhiteListPlayers string
	Ops              string
}

func getEnvVars() ServerProperties {
	return ServerProperties{
		MOTD:             os.Getenv("MOTD"),
		Memory:           os.Getenv("MEMORY"),
		Port:             os.Getenv("PORT"),
		Seed:             os.Getenv("SEED"),
		GameMode:         os.Getenv("GAMEMODE"),
		Difficulty:       os.Getenv("DIFFICULTY"),
		WhiteList:        os.Getenv("WHITELIST"),
		WhiteListPlayers: os.Getenv("WHITELIST_PLAYERS"),
		Ops:              os.Getenv("OPS"),
	}
}

const serverPropertiesTemplate = `motd={{.MOTD}}
server-port={{.Port}}
level-seed={{.Seed}}
gamemode={{.GameMode}}
difficulty={{.Difficulty}}
white-list={{.WhiteList}}
ops={{.Ops}}
`

func createServerPropertiesFile(props ServerProperties) {
	tmpl, err := template.New("server.properties").Parse(serverPropertiesTemplate)
    if err != nil {
        log.Fatalf("Error creating template: %v", err)
    }

    f, err := os.Create("server.properties")
    if err != nil {
        log.Fatalf("Error creating file: %v", err)
    }

    err = tmpl.Execute(f, props)
    if err != nil {
        log.Fatalf("Error executing template: %v", err)
    }
}

func main() {
    props := getEnvVars()
    if props.MOTD == "" {
    props.MOTD = "A Minecraft Server"
    }
    if props.Memory == "" {
    props.Memory = "1G"
    }
    if props.Port == "" {
    props.Port = "25565"
    }
    if props.Seed == "" {
    props.Seed = ""
    }
    if props.GameMode == "" {
    props.GameMode = "survival"
    }
    if props.Difficulty == "" {
    props.Difficulty = "easy"
    }
    if props.WhiteList == "" {
    props.WhiteList = "false"
    }
    if props.WhiteListPlayers == "" {
    props.WhiteListPlayers = ""
    }
    if props.Ops == "" {
    props.Ops = ""
    }
    createServerPropertiesFile(props)
}