package kubernetes

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type PodIpAndPort struct {
	podName   string
	nameSpace string
	ipAddress string
	ports     []string
}

func (w PodIpAndPort) String() string {
	return fmt.Sprintf("PodIpAndPort{podName: %v, nameSpace: %v, ipAddress: %v, ports: %v}", w.podName, w.nameSpace, w.ipAddress, w.ports)
}

func getIps() ([]PodIpAndPort, error) {
	jsonPath := "{range .items[*]}{.metadata.name}{\"\t\"}{.metadata.namespace}{\"\t\"}{.status.podIP}{\"\t\"}{range .spec.containers[*].ports[*]}{.containerPort}{\",\"}{end}" //{\"\n\"}{end}"
	cmd := exec.Command("kubectl", "get", "pods", "--selector=run=my-nginx", "-o", fmt.Sprintf("jsonpath=%s", jsonPath))

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error executing kubectl: %v", err)
		return nil, err
	}
	arr := strings.Split(string(out), ",")
	result := []PodIpAndPort{}
	for _, j := range arr {
		s := strings.Split(j, "\t")
		if j == "" {
			break
		}
		ports := strings.Split(s[3], ",")
		result = append(result, PodIpAndPort{
			podName:   s[0],
			nameSpace: s[1],
			ipAddress: s[2],
			ports:     ports,
		})
	}
	return result, nil
}
