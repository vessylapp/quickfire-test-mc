package main

import (
	"os"
	"os/exec"
	"text/template"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"io"
)

type Latest struct {
    Release string `json:"release"`
    Snapshot string `json:"snapshot"`
}

type Version struct {
    ID string `json:"id"`
    Type string `json:"type"`
    URL string `json:"url"`
    Time string `json:"time"`
    ReleaseTime string `json:"releaseTime"`
}

type VersionManifest struct {
    Latest Latest `json:"latest"`
    Versions []Version `json:"versions"`
}

type ServerProperties struct {
    Version          string
	MOTD             string
	Port             string
	Seed             string
	WhiteList        string
	Hardcore         string
	Rcon             string
	RconPassword     string
	RconPort         string
	MaxPlayers       string
}

func getEnvVars() ServerProperties {
	return ServerProperties{
	    Version:          os.Getenv("Version"),
		MOTD:             os.Getenv("MOTD"),
		Port:             os.Getenv("Port"),
		Seed:             os.Getenv("Seed"),
		WhiteList:        os.Getenv("WhiteList"),
		Hardcore:         os.Getenv("Hardcore"),
		Rcon:             os.Getenv("Rcon"),
		RconPassword:     os.Getenv("RconPassword"),
		RconPort:         os.Getenv("RconPort"),
		MaxPlayers:       os.Getenv("MaxPlayers"),
	}
}

const serverPropertiesTemplate = `motd={{.MOTD}}
server-port={{.Port}}
level-seed={{.Seed}}
white-list={{.WhiteList}}
hardcore={{.Hardcore}}
enable-rcon={{.Rcon}}
rcon.password={{.RconPassword}}
rcon.port={{.RconPort}}
max-players={{.MaxPlayers}}
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

func startServer(version string) {
    log.Println("Starting server...")
    cmd := exec.Command("java", "-jar", version + ".jar")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}

func installServerJar(version string) {
    log.Println("Downloading server.jar...")
    // Version manifest is in version_manifest.json, read it, find the version, and download the server.jar
    // Read the version_manifest.json
    file, err := os.Open("version_manifest.json")
    if err != nil {
        log.Fatalf("Error opening version_manifest.json: %v", err)
    }
    defer file.Close()

    // Read the file
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatalf("Error reading version_manifest.json: %v", err)
    }

    // Parse the JSON
    var manifest VersionManifest
    err = json.Unmarshal(data, &manifest)
    if err != nil {
        log.Fatalf("Error unmarshalling version_manifest.json: %v", err)
    }

    // Find the version
    var versionURL string
    for _, v := range manifest.Versions {
        if v.ID == version {
            versionURL = v.URL
            break
        }
    }

    if versionURL == "" {
        log.Fatalf("Version %s not found in version_manifest.json", version)
    }

    // make a request to the version URL
    resp, err := http.Get(versionURL)
    if err != nil {
        log.Fatalf("Error making request to version URL: %v", err)
    }
    defer resp.Body.Close()

    // create a file to write the server.jar to
    file, err = os.Create(version + ".jar")
    if err != nil {
        log.Fatalf("Error creating server.jar: %v", err)
    }
    defer file.Close()

    // the download the server.jar from body.downloads.server.url
    var versionData map[string]interface{}
    err = json.NewDecoder(resp.Body).Decode(&versionData)
    if err != nil {
        log.Fatalf("Error decoding version data: %v", err)
    }

    serverURL := versionData["downloads"].(map[string]interface{})["server"].(map[string]interface{})["url"].(string)

    resp, err = http.Get(serverURL)
    if err != nil {
        log.Fatalf("Error making request to server URL: %v", err)
    }
    defer resp.Body.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        log.Fatalf("Error copying server.jar: %v", err)
    }

    log.Println("server.jar downloaded")
}

func main() {
    props := getEnvVars()
    if props.Version == "" {
        log.Fatalf("Version is required")
    }
    if _, err := os.Stat("server.properties"); err == nil {
        log.Println("server.properties already exists, skipping creation")
        if _, err := os.Stat(props.Version + ".jar"); err == nil {
            log.Println("server.jar already exists, skipping download")
            startServer(props.Version)
        }
        return
    }
    if props.MOTD == "" {
    props.MOTD = "A Minecraft Server"
    }
    if props.Port == "" {
    props.Port = "25565"
    }
    if props.Seed == "" {
    props.Seed = ""
    }
    if props.WhiteList == "" {
    props.WhiteList = "false"
    }
    if props.Hardcore == "" {
    props.Hardcore = "false"
    }
    if props.Rcon == "" {
    props.Rcon = "false"
    }
    if props.RconPassword == "" {
    props.RconPassword = "password"
    }
    if props.RconPort == "" {
    props.RconPort = "25575"
    }
    if props.MaxPlayers == "" {
    props.MaxPlayers = "20"
    }
    createServerPropertiesFile(props)
    if _, err := os.Stat(props.Version + ".jar"); err == nil {
        log.Println("server.jar already exists, skipping download")
        startServer(props.Version)
    } else {
        installServerJar(props.Version)
        startServer(props.Version)
    }
}