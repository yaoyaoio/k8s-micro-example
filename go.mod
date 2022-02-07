module github.com/yaoliu/k8s-micro

go 1.14

require (
	github.com/ef-ds/deque v1.0.4-0.20190904040645-54cb57c252a1 // indirect
	github.com/evanphx/json-patch/v5 v5.0.0 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/golang/protobuf v1.3.5
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/micro/go-micro/v2 v2.3.0
	github.com/micro/go-plugins/broker/kafka/v2 v2.3.0 // indirect
	github.com/micro/go-plugins/config/source/configmap/v2 v2.3.0 // indirect
	github.com/micro/go-plugins/registry/kubernetes/v2 v2.3.0 // indirect
	github.com/nats-io/nats-server/v2 v2.2.0 // indirect
	go.etcd.io/bbolt v1.3.4 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.29.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => k8s.io/api v0.0.0-20190708174958-539a33f6e817
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
)
