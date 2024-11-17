//go:build !linux && !windows

package libnetwork

import (
	"errors"
	"net"
	"runtime"
)

func (c *Controller) getLBIndex(sid, nid string, ingressPorts []*PortConfig) int {
	return 0
}

func (c *Controller) cleanupServiceDiscovery(cleanupNID string) {}

func (c *Controller) cleanupServiceBindings(nid string) {}

func (c *Controller) addServiceBinding(svcName, svcID, nID, eID, containerName string, vip net.IP, ingressPorts []*PortConfig, serviceAliases, taskAliases []string, ip net.IP, method string) error {
	return errors.New("addServiceBinding is not supported on " + runtime.GOOS)
}

func (c *Controller) rmServiceBinding(svcName, svcID, nID, eID, containerName string, vip net.IP, ingressPorts []*PortConfig, serviceAliases []string, taskAliases []string, ip net.IP, method string, deleteSvcRecords bool, fullRemove bool) error {
	return errors.New("rmServiceBinding is not supported on " + runtime.GOOS)
}

func (c *Controller) addContainerNameResolution(nID, eID, containerName string, taskAliases []string, ip net.IP, method string) error {
	return errors.New("addContainerNameResolution is not supported on " + runtime.GOOS)
}

func (c *Controller) delContainerNameResolution(nID, eID, containerName string, taskAliases []string, ip net.IP, method string) error {
	return errors.New("delContainerNameResolution is not supported on " + runtime.GOOS)
}

func (sb *Sandbox) populateLoadBalancers(*Endpoint) {}

func arrangeIngressFilterRule() {}
