// Code generated by gotestmd DO NOT EDIT.
package afxdp

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/spire/single_cluster"
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
	r := s.Runner("../deployments-k8s/examples/afxdp")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete mutatingwebhookconfiguration nsm-mutating-webhook` + "\n" + `kubectl delete ns nsm-system`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/afxdp?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
}
func (s *Suite) TestKernel2IP2Kernel_ipv6() {
	r := s.Runner("../deployments-k8s/examples/features/ipv6/Kernel2IP2Kernel_ipv6")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel-ipv6`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/ipv6/Kernel2IP2Kernel_ipv6?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel-ipv6`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel-ipv6`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel-ipv6 -- ping -c 4 2001:db8::`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel-ipv6 -- ping -c 4 2001:db8::1`)
}
func (s *Suite) TestMemif2IP2Memif_ipv6() {
	r := s.Runner("../deployments-k8s/examples/features/ipv6/Memif2IP2Memif_ipv6")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ip2memif-ipv6`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/ipv6/Memif2IP2Memif_ipv6?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ip2memif-ipv6`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-memif2ip2memif-ipv6`)
	r.Run(`result=$(kubectl exec deployments/nsc-memif -n "ns-memif2ip2memif-ipv6" -- vppctl ping 2001:db8:: repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`result=$(kubectl exec deployments/nse-memif -n "ns-memif2ip2memif-ipv6" -- vppctl ping 2001:db8::1 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestKernel2Ethernet2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2Ethernet2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ethernet2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2Ethernet2Kernel?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ethernet2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ethernet2kernel`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ethernet2kernel -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ethernet2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2Ethernet2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2Ethernet2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ethernet2memif`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2Ethernet2Memif?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ethernet2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-kernel2ethernet2memif`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ethernet2memif -- ping -c 4 172.16.1.100`)
	r.Run(`result=$(kubectl exec deployments/nse-memif -n "ns-kernel2ethernet2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestKernel2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2IP2Kernel?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2memif`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2IP2Memif?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-kernel2ip2memif`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2memif -- ping -c 4 172.16.1.100`)
	r.Run(`result=$(kubectl exec deployments/nse-memif -n "ns-kernel2ip2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestMemif2Ethernet2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2Ethernet2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ethernet2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2Ethernet2Kernel?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ethernet2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-memif2ethernet2kernel`)
	r.Run(`result=$(kubectl exec deployments/nsc-memif -n "ns-memif2ethernet2kernel" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-memif2ethernet2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestMemif2Ethernet2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2Ethernet2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ethernet2memif`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2Ethernet2Memif?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ethernet2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-memif2ethernet2memif`)
	r.Run(`result=$(kubectl exec deployments/nsc-memif -n "ns-memif2ethernet2memif" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`result=$(kubectl exec deployments/nse-memif -n "ns-memif2ethernet2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestMemif2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ip2kernel`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2IP2Kernel?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ip2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-memif2ip2kernel`)
	r.Run(`result=$(kubectl exec deployments/nsc-memif -n "ns-memif2ip2kernel" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-memif2ip2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestMemif2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ip2memif`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2IP2Memif?ref=4381f9e51f314518878331b4c98da47596703766`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ip2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-memif2ip2memif`)
	r.Run(`result=$(kubectl exec deployments/nsc-memif -n "ns-memif2ip2memif" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`result=$(kubectl exec deployments/nse-memif -n "ns-memif2ip2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
