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
	r := s.Runner("../deployments-k8s/examples/interdomain/three_cluster_configuration/loadbalancer")
	s.T().Cleanup(func() {
		r.Run(`if [[ ! -z $CLUSTER1_CIDR ]]; then` + "\n" + `  kubectl --kubeconfig=$KUBECONFIG1 delete ns metallb-system  ` + "\n" + `fi`)
		r.Run(`if [[ ! -z $CLUSTER2_CIDR ]]; then` + "\n" + `  kubectl --kubeconfig=$KUBECONFIG2 delete ns metallb-system  ` + "\n" + `fi`)
		r.Run(`if [[ ! -z $CLUSTER3_CIDR ]]; then` + "\n" + `  kubectl --kubeconfig=$KUBECONFIG3 delete ns metallb-system  ` + "\n" + `fi`)
	})
	r.Run(`if [[ ! -z $CLUSTER1_CIDR ]]; then` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG1 apply -f - <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  namespace: metallb-system` + "\n" + `  name: config` + "\n" + `data:` + "\n" + `  config: |` + "\n" + `    address-pools:` + "\n" + `    - name: default` + "\n" + `      protocol: layer2` + "\n" + `      addresses:` + "\n" + `      - $CLUSTER1_CIDR` + "\n" + `EOF` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG1 wait --for=condition=ready --timeout=5m pod -l app=metallb -n metallb-system` + "\n" + `fi`)
	r.Run(`if [[ ! -z $CLUSTER2_CIDR ]]; then` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG2 apply -f - <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  namespace: metallb-system` + "\n" + `  name: config` + "\n" + `data:` + "\n" + `  config: |` + "\n" + `    address-pools:` + "\n" + `    - name: default` + "\n" + `      protocol: layer2` + "\n" + `      addresses:` + "\n" + `      - $CLUSTER2_CIDR` + "\n" + `EOF` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG2 wait --for=condition=ready --timeout=5m pod -l app=metallb -n metallb-system` + "\n" + `fi`)
	r.Run(`if [[ ! -z $CLUSTER3_CIDR ]]; then` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG3 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG3 apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG3 apply -f - <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  namespace: metallb-system` + "\n" + `  name: config` + "\n" + `data:` + "\n" + `  config: |` + "\n" + `    address-pools:` + "\n" + `    - name: default` + "\n" + `      protocol: layer2` + "\n" + `      addresses:` + "\n" + `      - $CLUSTER3_CIDR` + "\n" + `EOF` + "\n" + `    kubectl --kubeconfig=$KUBECONFIG3 wait --for=condition=ready --timeout=5m pod -l app=metallb -n metallb-system` + "\n" + `fi`)
}
func (s *Suite) Test() {}
