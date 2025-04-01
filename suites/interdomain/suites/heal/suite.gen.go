// Code generated by gotestmd DO NOT EDIT.
package heal

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/three_cluster_configuration/basic"
)

type Suite struct {
	base.Suite
	basicSuite basic.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.basicSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestHeal_floating_forwarder_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_floating-forwarder-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-forwarder-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-forwarder-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-forwarder-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-forwarder-death/cluster3?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-forwarder-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-forwarder-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-forwarder-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-forwarder-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-forwarder-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-forwarder-death -- ping -c 4 172.16.1.3`)
	r.Run(`LOCALFWD=$(kubectl --kubeconfig=$KUBECONFIG1 get pods -l app=forwarder-vpp -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `REMOTEFWD=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=forwarder-vpp -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete pod ${LOCALFWD} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete pod ${REMOTEFWD} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=1m pod -l app=forwarder-vpp -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=forwarder-vpp -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-forwarder-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-forwarder-death -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestHeal_floating_nse_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_floating-nse-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-nse-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-nse-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-nse-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nse-death/cluster3?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nse-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-nse-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nse-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-nse-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-nse-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-nse-death -- ping -c 4 172.16.1.3`)
	r.Run(`NSE=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=nse-kernel -n ns-floating-nse-death --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete pod ${NSE} -n ns-floating-nse-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-nse-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-nse-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-nse-death -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestHeal_floating_nsm_system_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_floating-nsm-system-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-floating-nsm-system-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-floating-nsm-system-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns ns-floating-nsm-system-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nsm-system-death/cluster3?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nsm-system-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-nsm-system-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_floating-nsm-system-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-nsm-system-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-nsm-system-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-nsm-system-death -- ping -c 4 172.16.1.3`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete mutatingwebhookconfiguration nsm-mutating-webhook` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete ns nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete mutatingwebhookconfiguration nsm-mutating-webhook` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete ns nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 delete ns nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/three_cluster_configuration/basic/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/three_cluster_configuration/basic/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/three_cluster_configuration/basic/cluster3?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 get services nsmgr-proxy -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'` + "\n" + `WH=$(kubectl --kubeconfig=$KUBECONFIG1 get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 get services nsmgr-proxy -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'` + "\n" + `WH=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG3 get services registry -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-floating-nsm-system-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-floating-nsm-system-death -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestHeal_interdomain_nsmgr_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_interdomain-nsmgr-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-interdomain-nsmgr-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-interdomain-nsmgr-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-nsmgr-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-nsmgr-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-nsmgr-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-nsmgr-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-nsmgr-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-nsmgr-death -- ping -c 4 172.16.1.3`)
	r.Run(`LOCALNSMGR=$(kubectl --kubeconfig=$KUBECONFIG1 get pods -l app=nsmgr -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `REMOTENSMGR=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=nsmgr -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete pod ${LOCALNSMGR} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete pod ${REMOTENSMGR} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=1m pod -l app=nsmgr -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nsmgr -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-nsmgr-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-nsmgr-death -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestHeal_interdomain_proxy_nsmgr_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_interdomain-proxy-nsmgr-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-interdomain-proxy-nsmgr-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-interdomain-proxy-nsmgr-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-proxy-nsmgr-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-proxy-nsmgr-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-proxy-nsmgr-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-proxy-nsmgr-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-proxy-nsmgr-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-proxy-nsmgr-death -- ping -c 4 172.16.1.3`)
	r.Run(`LOCAL_PROXY_NSMGR=$(kubectl --kubeconfig=$KUBECONFIG1 get pods -l app=nsmgr-proxy -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `REMOTE_PROXY_NSMGR=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=nsmgr-proxy -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete pod ${LOCAL_PROXY_NSMGR} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete pod ${REMOTE_PROXY_NSMGR} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=1m pod -l app=nsmgr-proxy -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nsmgr-proxy -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-proxy-nsmgr-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-proxy-nsmgr-death -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestHeal_interdomain_registry_death() {
	r := s.Runner("../deployments-k8s/examples/interdomain/usecases/heal_interdomain-registry-death")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG1 delete ns ns-interdomain-registry-death`)
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete ns ns-interdomain-registry-death`)
	})
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-registry-death/cluster2?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-registry-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/usecases/heal_interdomain-registry-death/cluster1?ref=807fdc7d8fe7279ace67218a626ae8fe2f8fe67d`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-registry-death`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-registry-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-registry-death -- ping -c 4 172.16.1.3`)
	r.Run(`REG=$(kubectl --kubeconfig=$KUBECONFIG2 get pods -l app=registry -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete pod ${REG} -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec pods/alpine -n ns-interdomain-registry-death -- ping -c 4 172.16.1.2`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 exec deployments/nse-kernel -n ns-interdomain-registry-death -- ping -c 4 172.16.1.3`)
}
