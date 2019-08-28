package list

import (
	"log"

	"github.com/gael-rozario/go-gcp/instances"
)

//Hostgrouplist ansible hostgroups map
type Hostgrouplist map[string][]string

//GetGcpAnsibleMap retries ansible mappings from gcp account
func GetGcpAnsibleMap(project string) Hostgrouplist {
	instancedata, err := instances.GetAllInstances(project)
	ansiblejsondata := make(Hostgrouplist)
	if err != nil {
		log.Fatal(err)
	}
	for _, instance := range instancedata.Items {
		if hostgroup, ok := instance.Labels["ansible_host"]; ok {
			iplist := ansiblejsondata[hostgroup]
			for _, netinter := range instance.NetworkInterfaces {
				iplist = append(iplist, netinter.NetworkIP)
			}
			ansiblejsondata[hostgroup] = iplist
		}
	}
	return ansiblejsondata
}
