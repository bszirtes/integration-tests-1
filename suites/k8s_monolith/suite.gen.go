// Code generated by gotestmd DO NOT EDIT.
package k8s_monolith

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/dns"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/docker"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/loadbalancer"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/spire"
)

type Suite struct {
	base.Suite
	dockerSuite       docker.Suite
	loadbalancerSuite loadbalancer.Suite
	dnsSuite          dns.Suite
	spireSuite        spire.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.dockerSuite, &s.loadbalancerSuite, &s.dnsSuite, &s.spireSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/k8s_monolith")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns nsm-system`)
	})
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/k8s_monolith/cluster-configuration?ref=ab453244ad6e6c66f0adc72c138d09ae46ed074b`)
	r.Run(`kubectl get services registry -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
}
func (s *Suite) TestKernel2Wireguard2Kernel() {
	r := s.Runner("../deployments-k8s/examples/k8s_monolith/usecases/Kernel2Wireguard2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ${NAMESPACE}`)
	})
	r.Run(`NAMESPACE=($(kubectl create -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ab453244ad6e6c66f0adc72c138d09ae46ed074b/examples/k8s_monolith/usecases/namespace.yaml)[0])` + "\n" + `NAMESPACE=${NAMESPACE:10}`)
	r.Run(`cat > kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ${NAMESPACE}` + "\n" + `` + "\n" + `bases:` + "\n" + `- https://github.com/networkservicemesh/deployments-k8s/apps/nsc-kernel?ref=ab453244ad6e6c66f0adc72c138d09ae46ed074b` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- patch-nsc.yaml` + "\n" + `EOF`)
	r.Run(`cat > patch-nsc.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nsc-kernel` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `        - name: nsc` + "\n" + `          env:` + "\n" + `            - name: NSM_NETWORK_SERVICES` + "\n" + `              value: kernel://docker-vl3/nsm-1` + "\n" + `EOF`)
	r.Run(`kubectl apply -k .`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel -n ${NAMESPACE}`)
	r.Run(`nscs=$(kubectl  get pods -l app=nsc-kernel -o go-template --template="{{range .items}}{{.metadata.name}} {{end}}" -n ${NAMESPACE})` + "\n" + `[[ ! -z $nscs ]]`)
	r.Run(`for nsc in $nscs` + "\n" + `do` + "\n" + `    ipAddr=$(kubectl exec -n ${NAMESPACE} $nsc -- ifconfig nsm-1)` + "\n" + `    ipAddr=$(echo $ipAddr | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `    for pinger in $nscs` + "\n" + `    do` + "\n" + `        echo $pinger pings $ipAddr` + "\n" + `        kubectl exec $pinger -n ${NAMESPACE} -- ping -c4 $ipAddr` + "\n" + `    done` + "\n" + `done`)
	r.Run(`for nsc in $nscs` + "\n" + `do` + "\n" + `    echo $nsc pings docker-nse` + "\n" + `    kubectl exec -n ${NAMESPACE} $nsc -- ping 169.254.0.1 -c4` + "\n" + `done`)
	r.Run(`for nsc in $nscs` + "\n" + `do` + "\n" + `    ipAddr=$(kubectl exec -n ${NAMESPACE} $nsc -- ifconfig nsm-1)` + "\n" + `    ipAddr=$(echo $ipAddr | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `    docker exec nse-simple-vl3-docker ping -c4 $ipAddr` + "\n" + `done`)
}
