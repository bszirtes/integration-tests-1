// Code generated by gotestmd DO NOT EDIT.
package single_cluster_csi

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
)

type Suite struct {
	base.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/spire/single_cluster_csi")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete crd clusterspiffeids.spire.spiffe.io` + "\n" + `kubectl delete crd clusterfederatedtrustdomains.spire.spiffe.io` + "\n" + `kubectl delete validatingwebhookconfiguration.admissionregistration.k8s.io/spire-controller-manager-webhook` + "\n" + `kubectl delete ns spire`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/spire/single_cluster_csi?ref=dba6fe311946b1c954583c5dbf0cf0cc4cfddf12`)
	r.Run(`kubectl wait -n spire --timeout=4m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/dba6fe311946b1c954583c5dbf0cf0cc4cfddf12/examples/spire/single_cluster/clusterspiffeid-template.yaml`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/dba6fe311946b1c954583c5dbf0cf0cc4cfddf12/examples/spire/base/clusterspiffeid-webhook-template.yaml`)
}
func (s *Suite) Test() {}
