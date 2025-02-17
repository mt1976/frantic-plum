package dockerhelpers

import (
	"log"
	"os"
	"strings"

	"github.com/mt1976/frantic-core/logHandler"
)

func DeployDefaultsPayload() error {
	list, err := os.ReadDir("startupPayload")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range list {
		if entry.IsDir() {
			continue
		}
		if strings.HasSuffix(entry.Name(), ".toml") {
			continue
		}
		if entry.Name() == ".DS_Store" {
			continue
		}
		if entry.Name() == ".keep" {
			continue
		}

		//logger.InfoLogger.Printf("Copying %v", entry.Name())
		from := "./defaults" + string(os.PathSeparator) + entry.Name()
		to := "./data" + string(os.PathSeparator) + "defaults" + string(os.PathSeparator) + entry.Name()
		logHandler.EventLogger.Printf("Copying [%v] to [%v]", from, to)
		err = startupCopyFile(from, to)
		if err != nil {
			logHandler.ErrorLogger.Println(err.Error())
		}
		//err = CopyFile("defaults/defaults.toml", "data/defaults/defaults.toml")

	}
	return err
}

func IsDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
