// Code generated by gotestmd DO NOT EDIT.
package loadbalancer

import (
	"github.com/stretchr/testify/suite"

	"github.com/bszirtes/integration-tests-1/extensions/base"
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
	r := s.Runner("../deployments-k8s/examples/k8s_monolith/configuration/loadbalancer")
	s.T().Cleanup(func() {
		r.Run(`if [[ ! -z $CLUSTER_CIDR ]]; then` + "\n" + `  kubectl delete ns metallb-system` + "\n" + `fi`)
	})
	r.Run(`if [[ ! -z $CLUSTER_CIDR ]]; then` + "\n" + `    kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml` + "\n" + `    kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml` + "\n" + `    kubectl apply -f - <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  namespace: metallb-system` + "\n" + `  name: config` + "\n" + `data:` + "\n" + `  config: |` + "\n" + `    address-pools:` + "\n" + `    - name: default` + "\n" + `      protocol: layer2` + "\n" + `      addresses:` + "\n" + `      - $CLUSTER_CIDR` + "\n" + `EOF` + "\n" + `    kubectl wait --for=condition=ready --timeout=5m pod -l app=metallb -n metallb-system` + "\n" + `fi`)
}
func (s *Suite) Test() {}
