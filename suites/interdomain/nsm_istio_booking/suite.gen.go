// Code generated by gotestmd DO NOT EDIT.
package nsm_istio_booking

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/dns"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/loadbalancer"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/nsm"
	"github.com/networkservicemesh/integration-tests/suites/interdomain/spire"
)

type Suite struct {
	base.Suite
	loadbalancerSuite loadbalancer.Suite
	dnsSuite          dns.Suite
	spireSuite        spire.Suite
	nsmSuite          nsm.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.loadbalancerSuite, &s.dnsSuite, &s.spireSuite, &s.nsmSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/interdomain/nsm_istio_booking")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/bookinfo/platform/kube/bookinfo.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_istio_booking/nse-auto-scale?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4 ` + "\n" + `kubectl --kubeconfig=$KUBECONFIG1 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ef02275eec2ad9601b41fabf40de496dd6b8dac4/examples/interdomain/nsm_istio_booking/productpage/productpage.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ef02275eec2ad9601b41fabf40de496dd6b8dac4/examples/interdomain/nsm_istio_booking/networkservice.yaml` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete ns istio-system` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 label namespace default istio-injection-` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete pods --all`)
	})
	r.Run(`curl -sL https://istio.io/downloadIstioctl | sh -` + "\n" + `export PATH=$PATH:$HOME/.istioctl/bin` + "\n" + `istioctl  install --set profile=minimal -y --kubeconfig=$KUBECONFIG2` + "\n" + `istioctl --kubeconfig=$KUBECONFIG2 proxy-status`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ef02275eec2ad9601b41fabf40de496dd6b8dac4/examples/interdomain/nsm_istio_booking/networkservice.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/ef02275eec2ad9601b41fabf40de496dd6b8dac4/examples/interdomain/nsm_istio_booking/productpage/productpage.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/networkservicemesh/deployments-k8s/examples/interdomain/nsm_istio_booking/nse-auto-scale?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 label namespace default istio-injection=enabled` + "\n" + `` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/istio/istio/release-1.13/samples/bookinfo/platform/kube/bookinfo.yaml`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 wait --timeout=2m --for=condition=ready pod -l app=productpage`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec deploy/productpage-v1 -c cmd-nsc -- apk add curl`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG1 exec deploy/productpage-v1 -c cmd-nsc -- curl -s productpage.default:9080/productpage | grep -o "<title>Simple Bookstore App</title>"`)
}
func (s *Suite) Test() {}
