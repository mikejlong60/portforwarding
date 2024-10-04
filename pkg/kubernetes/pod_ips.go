package kubernetes

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type PodIpAndPort struct {
	PodName   string
	NameSpace string
	IpAddress string
	Ports     []string
}

func (w PodIpAndPort) String() string {
	return fmt.Sprintf("PodIpAndPort{PodName: %v, NameSpace: %v, IpAddress: %v, Ports: %v}", w.PodName, w.NameSpace, w.IpAddress, w.Ports)
}

// iptables command to add rule for non-local traffic
// iptables -t nat -A PREROUTING -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080
// iptables command to add rule for local traffic
// iptables -t nat -A OUTPUT -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080
// iptables command to delete rule for non-local traffic
// iptables -t nat -D PREROUTING -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080

// iptables command for listing all rules for a given ip(i.e. 10.244.1.3)
//iptables -t nat -L -v -n | grep 10.244.1.3
//Note that you must issue the -D delete command, replacing -A with -D to remove te forwarding rule.

func GetIps() ([]PodIpAndPort, error) {
	jsonPath := "{range .items[*]}{.metadata.name}{\"\t\"}{.metadata.namespace}{\"\t\"}{.status.podIP}{\"\t\"}{range .spec.containers[*].ports[*]}{.containerPort}{\",\"}{end}" //{\"\n\"}{end}"
	cmd := exec.Command("kubectl", "get", "pods", "--selector=run=my-nginx", "-o", fmt.Sprintf("jsonpath=%s", jsonPath))
	piss := fmt.Sprintf("kubectl get pods --selector=run=my-nginx -o %v", fmt.Sprintf("jsonpath=%s", jsonPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error executing kubectl: %v", err)
		return nil, err
	}
	fmt.Println(piss)
	arr := strings.Split(string(out), ",")
	result := []PodIpAndPort{}
	for _, j := range arr {
		s := strings.Split(j, "\t")
		if j == "" {
			break
		}
		ports := strings.Split(s[3], ",")
		result = append(result, PodIpAndPort{
			PodName:   s[0],
			NameSpace: s[1],
			IpAddress: s[2],
			Ports:     ports,
		})
	}
	return result, nil
}
