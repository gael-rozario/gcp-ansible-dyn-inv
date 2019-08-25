package list

import (
	"encoding/json"
	"fmt"
)

type hostgroup map[string][]string

//List all hosts in inventory
func List() {
	kubernetesmaster := []string{"10.10.0.1", "10.10.0.2"}
	jsondata := make(hostgroup)
	jsondata["kubernetesmaster"] = kubernetesmaster

	jsonstring, _ := json.Marshal(jsondata)
	fmt.Println(string(jsonstring))
}
