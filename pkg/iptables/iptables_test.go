package iptables

import (
	"fmt"
	"github.com/mikejlong60/portforwarding/pkg/kubernetes"
	"testing"
)

func TestGetIps(t *testing.T) {
	pods, err := kubernetes.GetIps()
	if len(pods) != 5 {
		t.Errorf("Expected 5 pods. Actual number of pods:%v\n", len(pods))
	}
	err = AddOutputForwardingRule(pods[0])
	if err != nil {
		t.Errorf("Error adding output forwarding rule: %v", err)
	}
	fmt.Println(pods)
}
