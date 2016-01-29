package configurator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

type Config struct {
	DockerComposeName string   `json:"docker compose name"`
	Containers        []string `json:"container names"`
}

func StartContainers(c Config) error {

	for _, containerName := range c.Containers {
		containerProc := exec.Command(c.DockerComposeName, "up", "-d", containerName)

		fmt.Println(c.DockerComposeName, "up", "-d", containerName)
		containerProc.Start()
		err := containerProc.Wait()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func ReadConfig(fileName string) (Config, error) {
	config := Config{}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}