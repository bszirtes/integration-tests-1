// Code generated by gotestmd DO NOT EDIT.
package ovs

import (
	"github.com/stretchr/testify/suite"

	"github.com/bszirtes/integration-tests-1/extensions/base"
	"github.com/bszirtes/integration-tests-1/suites/spire/single_cluster"
)

type Suite struct {
	base.Suite
	single_clusterSuite single_cluster.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.single_clusterSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/ovs")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete mutatingwebhookconfiguration nsm-mutating-webhook` + "\n" + `kubectl delete ns nsm-system`)
	})
	r.Run(`kubectl apply -k https://github.com/bszirtes/deployments-k8s/examples/ovs?ref=v0.1.33`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
}
func (s *Suite) TestKernel2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/bszirtes/deployments-k8s/examples/use-cases/Kernel2IP2Kernel?ref=v0.1.33`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/bszirtes/deployments-k8s/examples/use-cases/Kernel2Kernel?ref=v0.1.33`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2kernel`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2kernel -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2KernelVLAN() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2KernelVLAN")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2kernel-vlan`)
	})
	r.Run(`kubectl apply -k https://github.com/bszirtes/deployments-k8s/examples/use-cases/Kernel2KernelVLAN?ref=v0.1.33`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel -n ns-kernel2kernel-vlan`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2kernel-vlan`)
	r.Run(`NSC=$((kubectl get pods -l app=nsc-kernel -n ns-kernel2kernel-vlan --template '{{range .items}}{{.metadata.name}}{{" "}}{{end}}') | cut -d' ' -f1)` + "\n" + `TARGET_IP=$(kubectl exec -ti ${NSC} -n ns-kernel2kernel-vlan -- ip route show | grep 172.16 | cut -d' ' -f1)`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-kernel2kernel-vlan --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec ${NSC} -n ns-kernel2kernel-vlan -- ping -c 4 ${TARGET_IP}`)
}
func (s *Suite) TestSmartVF2SmartVF() {
	r := s.Runner("../deployments-k8s/examples/use-cases/SmartVF2SmartVF")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-smartvf2smartvf`)
	})
	r.Run(`kubectl apply -k https://github.com/bszirtes/deployments-k8s/examples/use-cases/SmartVF2SmartVF?ref=v0.1.33`)
	r.Run(`kubectl -n ns-smartvf2smartvf wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel`)
	r.Run(`kubectl -n ns-smartvf2smartvf wait --for=condition=ready --timeout=1m pod -l app=nse-kernel`)
	r.Run(`kubectl -n ns-smartvf2smartvf exec deployments/nsc-kernel -- ping -c 4 172.16.1.100`)
}
