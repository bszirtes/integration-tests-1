// Code generated by gotestmd DO NOT EDIT.
package rvlanovs

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
	r := s.Runner("../deployments-k8s/examples/remotevlan/rvlanovs")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/remotevlan/rvlanovs?ref=463c38653dd7e3f190cf1b7c394f6daec8a07bde`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/remotevlan/rvlanovs?ref=463c38653dd7e3f190cf1b7c394f6daec8a07bde`)
	r.Run(`kubectl -n nsm-system wait --for=condition=ready --timeout=2m pod -l app=forwarder-ovs`)
}
func (s *Suite) TestKernel2RVlanBreakout() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2RVlanBreakout")
	s.T().Cleanup(func() {
		r.Run(`docker stop rvm-tester` + "\n" + `docker image rm rvm-tester:latest` + "\n" + `true`)
		r.Run(`kubectl delete ns ns-kernel2rvlan-breakout`)
	})
	r.Run(`kubectl create ns ns-kernel2rvlan-breakout`)
	r.Run(`cat > first-iperf-s.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: iperf1-s` + "\n" + `  labels:` + "\n" + `    app: iperf1-s` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: iperf1-s` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: iperf1-s` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://finance-bridge/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - iperf1-s` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: iperf-server` + "\n" + `        image: networkstatic/iperf3:latest` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        command: ["tail", "-f", "/dev/null"]` + "\n" + `EOF`)
	r.Run(`kubectl apply -n ns-kernel2rvlan-breakout -f ./first-iperf-s.yaml`)
	r.Run(`kubectl -n ns-kernel2rvlan-breakout wait --for=condition=ready --timeout=1m pod -l app=iperf1-s`)
	r.Run(`NSCS=($(kubectl get pods -l app=iperf1-s -n ns-kernel2rvlan-breakout --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))`)
	r.Run(`cat > Dockerfile <<EOF` + "\n" + `FROM networkstatic/iperf3` + "\n" + `` + "\n" + `RUN apt-get update \` + "\n" + `    && apt-get install -y ethtool iproute2 \` + "\n" + `    && rm -rf /var/lib/apt/lists/*` + "\n" + `` + "\n" + `ENTRYPOINT [ "tail", "-f", "/dev/null" ]` + "\n" + `EOF` + "\n" + `docker build . -t rvm-tester`)
	r.Run(`docker run --cap-add=NET_ADMIN --rm -d --network bridge-2 --name rvm-tester rvm-tester tail -f /dev/null` + "\n" + `docker exec rvm-tester ip link set eth0 down` + "\n" + `docker exec rvm-tester ip link add link eth0 name eth0.100 type vlan id 100` + "\n" + `docker exec rvm-tester ip link set eth0 up` + "\n" + `docker exec rvm-tester ip addr add 172.10.0.254/24 dev eth0.100` + "\n" + `docker exec rvm-tester ethtool -K eth0 tx off`)
	r.Run(`status=0` + "\n" + `    for nsc in "${NSCS[@]}"` + "\n" + `    do` + "\n" + `      IP_ADDRESS=$(kubectl exec ${nsc} -c cmd-nsc -n ns-kernel2rvlan-breakout -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `      kubectl exec ${nsc} -c iperf-server -n ns-kernel2rvlan-breakout -- iperf3 -sD -B ${IP_ADDRESS} -1` + "\n" + `      docker exec rvm-tester iperf3 -i0 -t 25 -c ${IP_ADDRESS}` + "\n" + `      if test $? -ne 0` + "\n" + `      then` + "\n" + `        status=1` + "\n" + `      fi` + "\n" + `    done` + "\n" + `    if test ${status} -eq 1` + "\n" + `    then` + "\n" + `      false` + "\n" + `    fi`)
	r.Run(`status=0` + "\n" + `    for nsc in "${NSCS[@]}"` + "\n" + `    do` + "\n" + `      IP_ADDRESS=$(kubectl exec ${nsc} -c cmd-nsc -n ns-kernel2rvlan-breakout -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `      kubectl exec ${nsc} -c iperf-server -n ns-kernel2rvlan-breakout -- iperf3 -sD -B ${IP_ADDRESS} -1` + "\n" + `      docker exec rvm-tester iperf3 -i0 -t 5 -u -c ${IP_ADDRESS}` + "\n" + `      if test $? -ne 0` + "\n" + `      then` + "\n" + `        status=1` + "\n" + `      fi` + "\n" + `    done` + "\n" + `    if test ${status} -eq 1` + "\n" + `    then` + "\n" + `      false` + "\n" + `    fi`)
	r.Run(`status=0` + "\n" + `    for nsc in "${NSCS[@]}"` + "\n" + `    do` + "\n" + `      docker exec rvm-tester iperf3 -sD -B 172.10.0.254 -1` + "\n" + `      kubectl exec ${nsc} -c iperf-server -n ns-kernel2rvlan-breakout -- iperf3 -i0 -t 5 -c 172.10.0.254` + "\n" + `      if test $? -ne 0` + "\n" + `      then` + "\n" + `        status=1` + "\n" + `      fi` + "\n" + `    done` + "\n" + `    if test ${status} -eq 1` + "\n" + `    then` + "\n" + `      false` + "\n" + `    fi`)
	r.Run(`status=0` + "\n" + `    for nsc in "${NSCS[@]}"` + "\n" + `    do` + "\n" + `      docker exec rvm-tester iperf3 -sD -B 172.10.0.254 -1` + "\n" + `      kubectl exec ${NSCS[1]} -c iperf-server -n ns-kernel2rvlan-breakout -- iperf3 -i0 -t 5 -u -c 172.10.0.254` + "\n" + `      if test $? -ne 0` + "\n" + `      then` + "\n" + `        status=1` + "\n" + `      fi` + "\n" + `    done` + "\n" + `    if test ${status} -eq 1` + "\n" + `    then` + "\n" + `      false` + "\n" + `    fi`)
}
func (s *Suite) TestKernel2RVlanInternal() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2RVlanInternal")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2rvlan-internal`)
	})
	r.Run(`kubectl create ns ns-kernel2rvlan-internal`)
	r.Run(`cat > first-iperf-s.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: iperf1-s` + "\n" + `  labels:` + "\n" + `    app: iperf1-s` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: iperf1-s` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: iperf1-s` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://finance-bridge/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - iperf1-s` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: iperf-server` + "\n" + `        image: networkstatic/iperf3:latest` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        command: ["tail", "-f", "/dev/null"]` + "\n" + `EOF`)
	r.Run(`cat > kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ns-kernel2rvlan-internal` + "\n" + `` + "\n" + `resources:` + "\n" + `- first-iperf-s.yaml` + "\n" + `` + "\n" + `EOF`)
	r.Run(`kubectl apply -k .`)
	r.Run(`kubectl -n ns-kernel2rvlan-internal wait --for=condition=ready --timeout=1m pod -l app=iperf1-s`)
	r.Run(`NSCS=($(kubectl get pods -l app=iperf1-s -n ns-kernel2rvlan-internal --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))`)
	r.Run(`IP_ADDR=$(kubectl exec ${NSCS[0]} -c cmd-nsc -n ns-kernel2rvlan-internal -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `    kubectl exec ${NSCS[0]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -sD -B ${IP_ADDR} -1` + "\n" + `    kubectl exec ${NSCS[1]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -i0 -t 5 -c ${IP_ADDR}`)
	r.Run(`IP_ADDR=$(kubectl exec ${NSCS[1]} -c cmd-nsc -n ns-kernel2rvlan-internal -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `    kubectl exec ${NSCS[1]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -sD -B ${IP_ADDR} -1` + "\n" + `    kubectl exec ${NSCS[0]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -i0 -t 5 -u -c ${IP_ADDR}`)
	r.Run(`IP_ADDR=$(kubectl exec ${NSCS[0]} -c cmd-nsc -n ns-kernel2rvlan-internal -- ip -6 a s nsm-1 scope global | grep -oP '(?<=inet6\s)([0-9a-f:]+:+)+[0-9a-f]+')` + "\n" + `    kubectl exec ${NSCS[0]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -sD -B ${IP_ADDR} -1` + "\n" + `    kubectl exec ${NSCS[1]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -i0 -t 5 -6 -c ${IP_ADDR}`)
	r.Run(`IP_ADDR=$(kubectl exec ${NSCS[1]} -c cmd-nsc -n ns-kernel2rvlan-internal -- ip -6 a s nsm-1 scope global | grep -oP '(?<=inet6\s)([0-9a-f:]+:+)+[0-9a-f]+')` + "\n" + `    kubectl exec ${NSCS[1]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -sD -B ${IP_ADDR} -1` + "\n" + `    kubectl exec ${NSCS[0]} -c iperf-server -n ns-kernel2rvlan-internal -- iperf3 -i0 -t 5 -6 -u -c ${IP_ADDR}`)
}
func (s *Suite) TestKernel2RVlanMultiNS() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2RVlanMultiNS")
	s.T().Cleanup(func() {
		r.Run(`docker stop rvm-tester && \` + "\n" + `docker image rm rvm-tester:latest` + "\n" + `true`)
		r.Run(`kubectl delete --namespace=nsm-system -f client.yaml`)
		r.Run(`kubectl delete ns ns-kernel2vlan-multins-1`)
		r.Run(`kubectl delete ns ns-kernel2vlan-multins-2`)
		r.Run(`rm -rf ns-1 ns-2`)
	})
	r.Run(`kubectl create ns ns-kernel2vlan-multins-1` + "\n" + `kubectl create ns ns-kernel2vlan-multins-2`)
	r.Run(`mkdir -p ns-1 ns-2`)
	r.Run(`cat > ns-1/first-client.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: alpine-1` + "\n" + `  labels:` + "\n" + `    app: alpine-1` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: alpine-1` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: alpine-1` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://private-bridge.ns-kernel2vlan-multins-1/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - alpine-1` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: alpine` + "\n" + `        image: alpine:3.15.0` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        stdin: true` + "\n" + `        tty: true` + "\n" + `EOF`)
	r.Run(`cat > ns-1/patch-nse.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nse-remote-vlan` + "\n" + `spec:` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `      - name: nse` + "\n" + `        env:` + "\n" + `          - name: NSM_CONNECT_TO` + "\n" + `            value: "registry.nsm-system:5002"` + "\n" + `          - name: NSM_SERVICES` + "\n" + `            value: "private-bridge.ns-kernel2vlan-multins-1 { vlan: 0; via: gw1 }"` + "\n" + `          - name: NSM_CIDR_PREFIX` + "\n" + `            value: 172.10.1.0/24` + "\n" + `EOF`)
	r.Run(`cat > ns-1/kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ns-kernel2vlan-multins-1` + "\n" + `` + "\n" + `resources:` + "\n" + `- first-client.yaml` + "\n" + `` + "\n" + `bases:` + "\n" + `- ../../../../apps/nse-remote-vlan` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- patch-nse.yaml` + "\n" + `EOF`)
	r.Run(`kubectl apply -k ./ns-1`)
	r.Run(`cat > ns-2/second-client.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: alpine-2` + "\n" + `  labels:` + "\n" + `    app: alpine-2` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: alpine-2` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: alpine-2` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://blue-bridge.ns-kernel2vlan-multins-2/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - alpine-2` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: alpine` + "\n" + `        image: alpine:3.15.0` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        stdin: true` + "\n" + `        tty: true` + "\n" + `EOF` + "\n" + `cat > ns-2/third-client.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: alpine-3` + "\n" + `  labels:` + "\n" + `    app: alpine-3` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: alpine-3` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: alpine-3` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://green-bridge.ns-kernel2vlan-multins-2/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - alpine-3` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: alpine` + "\n" + `        image: alpine:3.15.0` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        stdin: true` + "\n" + `        tty: true` + "\n" + `EOF`)
	r.Run(`cat > ns-2/patch-nse.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: nse-remote-vlan` + "\n" + `spec:` + "\n" + `  template:` + "\n" + `    spec:` + "\n" + `      containers:` + "\n" + `      - name: nse` + "\n" + `        env:` + "\n" + `          - name: NSM_CONNECT_TO` + "\n" + `            value: "registry.nsm-system:5002"` + "\n" + `          - name: NSM_SERVICES` + "\n" + `            value: "blue-bridge.ns-kernel2vlan-multins-2 { vlan: 300; via: gw1 }, green-bridge.ns-kernel2vlan-multins-2 { vlan: 400; via: gw1 }"` + "\n" + `          - name: NSM_CIDR_PREFIX` + "\n" + `            value: 172.10.2.0/24` + "\n" + `EOF`)
	r.Run(`cat > ns-2/kustomization.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: kustomize.config.k8s.io/v1beta1` + "\n" + `kind: Kustomization` + "\n" + `` + "\n" + `namespace: ns-kernel2vlan-multins-2` + "\n" + `` + "\n" + `resources:` + "\n" + `- second-client.yaml` + "\n" + `- third-client.yaml` + "\n" + `` + "\n" + `bases:` + "\n" + `- https://github.com/networkservicemesh/deployments-k8s/apps/nse-remote-vlan?ref=463c38653dd7e3f190cf1b7c394f6daec8a07bde` + "\n" + `` + "\n" + `nameSuffix: -bg` + "\n" + `` + "\n" + `patchesStrategicMerge:` + "\n" + `- patch-nse.yaml` + "\n" + `EOF`)
	r.Run(`kubectl apply -k ./ns-2`)
	r.Run(`cat > client.yaml <<EOF` + "\n" + `---` + "\n" + `apiVersion: apps/v1` + "\n" + `kind: Deployment` + "\n" + `metadata:` + "\n" + `  name: alpine-4` + "\n" + `  labels:` + "\n" + `    app: alpine-4` + "\n" + `spec:` + "\n" + `  replicas: 2` + "\n" + `  selector:` + "\n" + `    matchLabels:` + "\n" + `      app: alpine-4` + "\n" + `  template:` + "\n" + `    metadata:` + "\n" + `      labels:` + "\n" + `        app: alpine-4` + "\n" + `      annotations:` + "\n" + `        networkservicemesh.io: kernel://finance-bridge/nsm-1` + "\n" + `    spec:` + "\n" + `      affinity:` + "\n" + `        podAntiAffinity:` + "\n" + `          requiredDuringSchedulingIgnoredDuringExecution:` + "\n" + `          - labelSelector:` + "\n" + `              matchExpressions:` + "\n" + `              - key: app` + "\n" + `                operator: In` + "\n" + `                values:` + "\n" + `                - alpine-4` + "\n" + `            topologyKey: "kubernetes.io/hostname"` + "\n" + `      containers:` + "\n" + `      - name: alpine` + "\n" + `        image: alpine:3.15.0` + "\n" + `        imagePullPolicy: IfNotPresent` + "\n" + `        stdin: true` + "\n" + `        tty: true` + "\n" + `EOF`)
	r.Run(`kubectl apply -n nsm-system -f client.yaml`)
	r.Run(`kubectl -n ns-kernel2vlan-multins-1 wait --for=condition=ready --timeout=1m pod -l app=nse-remote-vlan`)
	r.Run(`kubectl -n ns-kernel2vlan-multins-1 wait --for=condition=ready --timeout=1m pod -l app=alpine-1`)
	r.Run(`kubectl -n ns-kernel2vlan-multins-2 wait --for=condition=ready --timeout=1m pod -l app=nse-remote-vlan`)
	r.Run(`kubectl -n ns-kernel2vlan-multins-2 wait --for=condition=ready --timeout=1m pod -l app=alpine-2`)
	r.Run(`kubectl -n ns-kernel2vlan-multins-2 wait --for=condition=ready --timeout=1m pod -l app=alpine-3`)
	r.Run(`kubectl -n nsm-system wait --for=condition=ready --timeout=1m pod -l app=alpine-4`)
	r.Run(`cat > Dockerfile <<EOF` + "\n" + `FROM alpine:3.15.0` + "\n" + `` + "\n" + `RUN apk add ethtool` + "\n" + `` + "\n" + `ENTRYPOINT [ "tail", "-f", "/dev/null" ]` + "\n" + `EOF` + "\n" + `docker build . -t rvm-tester`)
	r.Run(`docker run --cap-add=NET_ADMIN --rm -d --network bridge-2 --name rvm-tester rvm-tester tail -f /dev/null` + "\n" + `docker exec rvm-tester ip link set eth0 down` + "\n" + `docker exec rvm-tester ip link add link eth0 name eth0.100 type vlan id 100` + "\n" + `docker exec rvm-tester ip link add link eth0 name eth0.300 type vlan id 300` + "\n" + `docker exec rvm-tester ip link add link eth0 name eth0.400 type vlan id 400` + "\n" + `docker exec rvm-tester ip link set eth0 up` + "\n" + `docker exec rvm-tester ip addr add 172.10.0.254/24 dev eth0.100` + "\n" + `docker exec rvm-tester ip addr add 172.10.1.254/24 dev eth0` + "\n" + `docker exec rvm-tester ip addr add 172.10.2.254/24 dev eth0.300` + "\n" + `docker exec rvm-tester ip addr add 172.10.2.253/24 dev eth0.400` + "\n" + `docker exec rvm-tester ethtool -K eth0 tx off`)
	r.Run(`NSCS=($(kubectl get pods -l app=alpine-1 -n ns-kernel2vlan-multins-1 --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))`)
	r.Run(`declare -A IP_ADDR` + "\n" + `for nsc in "${NSCS[@]}"` + "\n" + `do` + "\n" + `  IP_ADDR[$nsc]=$(kubectl exec ${nsc} -n ns-kernel2vlan-multins-1 -c alpine -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `done`)
	r.Run(`status=0` + "\n" + `for nsc in "${NSCS[@]}"` + "\n" + `do` + "\n" + `  for vlan_if_name in eth0.100 eth0.300 eth0.400` + "\n" + `  do` + "\n" + `    docker exec rvm-tester ping -w 1 -c 1 ${IP_ADDR[$nsc]} -I ${vlan_if_name}` + "\n" + `    if test $? -eq 0` + "\n" + `      then` + "\n" + `        status=2` + "\n" + `    fi` + "\n" + `  done` + "\n" + `  docker exec rvm-tester ping -c 1 ${IP_ADDR[$nsc]} -I eth0` + "\n" + `  if test $? -ne 0` + "\n" + `    then` + "\n" + `      status=1` + "\n" + `  fi` + "\n" + `done` + "\n" + `if test ${status} -eq 1` + "\n" + `  then` + "\n" + `    false` + "\n" + `fi`)
	r.Run(`NSCS_BLUE=($(kubectl get pods -l app=alpine-2 -n ns-kernel2vlan-multins-2 --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))` + "\n" + `NSCS_GREEN=($(kubectl get pods -l app=alpine-3 -n ns-kernel2vlan-multins-2 --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))`)
	r.Run(`declare -A IP_ADDR_BLUE` + "\n" + `for nsc in "${NSCS_BLUE[@]}"` + "\n" + `do` + "\n" + `  IP_ADDR_BLUE[$nsc]=$(kubectl exec ${nsc} -n ns-kernel2vlan-multins-2 -c alpine -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `done` + "\n" + `declare -A IP_ADDR_GREEN` + "\n" + `for nsc in "${NSCS_GREEN[@]}"` + "\n" + `do` + "\n" + `  IP_ADDR_GREEN[$nsc]=$(kubectl exec ${nsc} -n ns-kernel2vlan-multins-2 -c alpine -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `done`)
	r.Run(`status=0` + "\n" + `for nsc in "${NSCS_BLUE[@]}"` + "\n" + `do` + "\n" + `  for vlan_if_name in eth0.100 eth0 eth0.400` + "\n" + `  do` + "\n" + `    docker exec rvm-tester ping -w 1 -c 1 ${IP_ADDR_BLUE[$nsc]} -I ${vlan_if_name}` + "\n" + `    if test $? -eq 0` + "\n" + `      then` + "\n" + `        status=2` + "\n" + `    fi` + "\n" + `  done` + "\n" + `  docker exec rvm-tester ping -c 1 ${IP_ADDR_BLUE[$nsc]} -I eth0.300` + "\n" + `  if test $? -ne 0` + "\n" + `    then` + "\n" + `      status=1` + "\n" + `  fi` + "\n" + `done` + "\n" + `for nsc in "${NSCS_GREEN[@]}"` + "\n" + `do` + "\n" + `  for vlan_if_name in eth0.100 eth0 eth0.300` + "\n" + `  do` + "\n" + `    docker exec rvm-tester ping -w 1 -c 1 ${IP_ADDR_GREEN[$nsc]} -I ${vlan_if_name}` + "\n" + `    if test $? -eq 0` + "\n" + `      then` + "\n" + `        status=2` + "\n" + `    fi` + "\n" + `  done` + "\n" + `  docker exec rvm-tester ping -c 1 ${IP_ADDR_GREEN[$nsc]} -I eth0.400` + "\n" + `  if test $? -ne 0` + "\n" + `    then` + "\n" + `      status=1` + "\n" + `  fi` + "\n" + `done` + "\n" + `if test ${status} -eq 1` + "\n" + `  then` + "\n" + `    false` + "\n" + `fi`)
	r.Run(`NSCS=($(kubectl get pods -l app=alpine-4 -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}'))`)
	r.Run(`declare -A IP_ADDR` + "\n" + `for nsc in "${NSCS[@]}"` + "\n" + `do` + "\n" + `  IP_ADDR[$nsc]=$(kubectl exec ${nsc} -n nsm-system -c alpine -- ip -4 addr show nsm-1 | grep -oP '(?<=inet\s)\d+(\.\d+){3}')` + "\n" + `done`)
	r.Run(`status=0` + "\n" + `for nsc in "${NSCS[@]}"` + "\n" + `do` + "\n" + `  for vlan_if_name in eth0 eth0.300 eth0.400` + "\n" + `  do` + "\n" + `    docker exec rvm-tester ping -w 1 -c 1 ${IP_ADDR[$nsc]} -I ${vlan_if_name}` + "\n" + `    if test $? -eq 0` + "\n" + `      then` + "\n" + `        status=2` + "\n" + `    fi` + "\n" + `  done` + "\n" + `  docker exec rvm-tester ping -c 1 ${IP_ADDR[$nsc]} -I eth0.100` + "\n" + `  if test $? -ne 0` + "\n" + `    then` + "\n" + `      status=1` + "\n" + `  fi` + "\n" + `done` + "\n" + `if test ${status} -eq 1` + "\n" + `  then` + "\n" + `    false` + "\n" + `fi`)
}
