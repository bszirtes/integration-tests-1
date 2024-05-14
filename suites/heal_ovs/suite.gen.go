// Code generated by gotestmd DO NOT EDIT.
package heal_ovs

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/ovs"
)

type Suite struct {
	base.Suite
	ovsSuite ovs.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.ovsSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestLocal_forwarder_death() {
	r := s.Runner("../deployments-k8s/examples/heal/local-forwarder-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-local-forwarder-death`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/local-forwarder-death?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-local-forwarder-death`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-local-forwarder-death`)
	r.Run(`kubectl exec pods/alpine -n ns-local-forwarder-death -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-forwarder-death -- ping -c 4 172.16.1.101`)
	r.Run(`NSC_NODE=$(kubectl get pods -l app=alpine -n ns-local-forwarder-death --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSC_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod -n nsm-system ${FORWARDER}`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSC_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-local-forwarder-death -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-forwarder-death -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestLocal_forwarder_remote_forwarder() {
	r := s.Runner("../deployments-k8s/examples/heal/local-forwarder-remote-forwarder")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-local-forwarder-remote-forwarder`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/local-forwarder-remote-forwarder?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-local-forwarder-remote-forwarder`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-local-forwarder-remote-forwarder`)
	r.Run(`kubectl exec pods/alpine -n ns-local-forwarder-remote-forwarder -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-forwarder-remote-forwarder -- ping -c 4 172.16.1.101`)
	r.Run(`NSC_NODE=$(kubectl get pods -l app=alpine -n ns-local-forwarder-remote-forwarder --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')` + "\n" + `NSE_NODE=$(kubectl get pods -l app=nse-kernel -n ns-local-forwarder-remote-forwarder --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER1=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSC_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER2=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod ${FORWARDER1} -n nsm-system`)
	r.Run(`kubectl delete pod ${FORWARDER2} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSC_NODE} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-local-forwarder-remote-forwarder -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-forwarder-remote-forwarder -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestLocal_nsmgr_restart() {
	r := s.Runner("../deployments-k8s/examples/heal/local-nsmgr-restart")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-local-nsmgr-restart`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/local-nsmgr-restart?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-local-nsmgr-restart`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-local-nsmgr-restart`)
	r.Run(`kubectl exec pods/alpine -n ns-local-nsmgr-restart -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-nsmgr-restart -- ping -c 4 172.16.1.101`)
	r.Run(`NSC_NODE=$(kubectl get pods -l app=alpine -n ns-local-nsmgr-restart --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`NSMGR=$(kubectl get pods -l app=nsmgr --field-selector spec.nodeName==${NSC_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod ${NSMGR} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsmgr --field-selector spec.nodeName==${NSC_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-local-nsmgr-restart -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-local-nsmgr-restart -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestRegistry_remote_forwarder() {
	r := s.Runner("../deployments-k8s/examples/heal/registry-remote-forwarder")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-registry-remote-forwarder`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/registry-remote-forwarder?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-registry-remote-forwarder`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-registry-remote-forwarder`)
	r.Run(`kubectl exec pods/alpine -n ns-registry-remote-forwarder -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-registry-remote-forwarder -- ping -c 4 172.16.1.101`)
	r.Run(`NSE_NODE=$(kubectl get pods -l app=nse-kernel -n ns-registry-remote-forwarder --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`REGISTRY=$(kubectl get pods -l app=registry -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod ${REGISTRY} -n nsm-system`)
	r.Run(`kubectl delete pod ${FORWARDER} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-registry-remote-forwarder -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-registry-remote-forwarder -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestRegistry_restart() {
	r := s.Runner("../deployments-k8s/examples/heal/registry-restart")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-registry-restart`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/registry-restart/registry-before-death?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-registry-restart`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-registry-restart`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-registry-restart --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-registry-restart --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec pods/alpine -n ns-registry-restart -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-registry-restart -- ping -c 4 172.16.1.101`)
	r.Run(`REGISTRY=$(kubectl get pods -l app=registry -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod ${REGISTRY} -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/registry-restart/registry-after-death?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine-new -n ns-registry-restart`)
	r.Run(`kubectl exec pods/alpine-new -n ns-registry-restart -- ping -c 4 172.16.1.102`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-registry-restart -- ping -c 4 172.16.1.103`)
}
func (s *Suite) TestRemote_forwarder_death() {
	r := s.Runner("../deployments-k8s/examples/heal/remote-forwarder-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-remote-forwarder-death`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/remote-forwarder-death?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-remote-forwarder-death`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-remote-forwarder-death`)
	r.Run(`kubectl exec pods/alpine -n ns-remote-forwarder-death -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-remote-forwarder-death -- ping -c 4 172.16.1.101`)
	r.Run(`NSE_NODE=$(kubectl get pods -l app=nse-kernel -n ns-remote-forwarder-death --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod -n nsm-system ${FORWARDER}`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-remote-forwarder-death -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-remote-forwarder-death -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestRemote_forwarder_death_ip() {
	r := s.Runner("../deployments-k8s/examples/heal/remote-forwarder-death-ip")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-remote-forwarder-death-ip`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/heal/remote-forwarder-death-ip?ref=ac1b4ec4d829bcd44944db761cc15ede29827052`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-remote-forwarder-death-ip`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-remote-forwarder-death-ip`)
	r.Run(`kubectl exec pods/alpine -n ns-remote-forwarder-death-ip -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-remote-forwarder-death-ip -- ping -c 4 172.16.1.101`)
	r.Run(`NSE_NODE=$(kubectl get pods -l app=nse-kernel -n ns-remote-forwarder-death-ip --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}')`)
	r.Run(`FORWARDER=$(kubectl get pods -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl delete pod -n nsm-system ${FORWARDER}`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l 'app in (forwarder-ovs, forwarder-vpp)' --field-selector spec.nodeName==${NSE_NODE} -n nsm-system`)
	r.Run(`kubectl exec pods/alpine -n ns-remote-forwarder-death-ip -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-remote-forwarder-death-ip -- ping -c 4 172.16.1.101`)
}
