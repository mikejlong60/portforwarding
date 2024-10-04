package iptables

import (
	"fmt"
	"github.com/mikejlong60/portforwarding/pkg/kubernetes"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

// iptables command to add rule for non-local traffic
// iptables -t nat -A PREROUTING -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080
// iptables command to add rule for local traffic
// iptables -t nat -A OUTPUT -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080
// iptables command to delete rule for non-local traffic
// iptables -t nat -D PREROUTING -d 10.244.1.3 -p tcp --dport 80 -j REDIRECT --to-port 8080

// iptables command for listing all rules for a given ip(i.e. 10.244.1.3)
//iptables -t nat -L -v -n | grep 10.244.1.3
//Note that you must issue the -D delete command, replacing -A with -D to remove te forwarding rule.

func executeRule(cmd *exec.Cmd) error {
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error executing kubectl: %v", err)
		return err
	}
	fmt.Println(string(out))
	return nil
}

func AddOutputForwardingRule(p kubernetes.PodIpAndPort) error {
	//addForwardCmd := fmt.Sprintf("iptables -t nat -A OUTPUT -d %v -p tcp --dport %v -j REDIRECT --to-port 8080", p.IpAddress, p.Ports[0])

	piss := fmt.Sprintf("iptables -t nat -A OUTPUT  %v -p tcp %v -j REDIRECT --to-port 8080", fmt.Sprintf("-d %v", p.IpAddress), fmt.Sprintf("--dport %v", p.Ports[0]))
	fmt.Println(piss)
	//cmd := exec.Command("sudo", "iptables", "-t nat", "-A OUTPUT", fmt.Sprintf("-d %v", p.IpAddress), "-p tcp", fmt.Sprintf("--dport %v", p.Ports[0]), "-j REDIRECT", "--to-port 8080")
	cmd := exec.Command("sudo", "whoami")
	err := executeRule(cmd)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOutputForwardingRule(p kubernetes.PodIpAndPort) error {
	addForwardCmd := fmt.Sprintf("iptables -t nat -D OUTPUT -d %v -p tcp --dport %v -j REDIRECT --to-port 8080", p.IpAddress, p.Ports[0])

	cmd := exec.Command(addForwardCmd)
	err := executeRule(cmd)
	if err != nil {
		return err
	}
	return nil
}

func DeletePreroutingForwardingRule(p kubernetes.PodIpAndPort) error {
	//addForwardCmd := fmt.Sprintf("iptables -t nat -D PREROUTING -d %v -p tcp --dport %v -j REDIRECT --to-port 8080", p.IpAddress, p.Ports[0])
	fromPort := fmt.Sprintf("--dport %v", p.Ports[0])
	prerouting := fmt.Sprintf("-D PREROUTING %v", p.IpAddress)
	//addForwardCmd := fmt.Sprintf("iptables -t nat -D PREROUTING -d %v -p tcp --dport %v -j REDIRECT --to-port 8080", p.IpAddress, p.Ports[0])
	cmd := exec.Command("iptables", "-t nat", prerouting, "-p tcp", fromPort, "-j REDIRECT", "--to-port 8080")
	err := executeRule(cmd)
	if err != nil {
		return err
	}
	return nil
}
