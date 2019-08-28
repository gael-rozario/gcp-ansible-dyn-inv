package list

import (
	"encoding/json"
	"fmt"
)

//List all hosts in inventory
func List() {
	gcpproject := "angelic-bond-246708"
	ansiblejson := GetGcpAnsibleMap(gcpproject)
	jsonstring, _ := json.Marshal(ansiblejson)
	fmt.Println(string(jsonstring))
}
