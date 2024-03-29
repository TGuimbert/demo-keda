# demo-keda

```shell
kind create cluster
k apply --server-side -f keda-2.13.1.yaml
k apply -f kafka.yaml
k logs -n kafka kafka-0
k exec -it -n kafka kafka-0 -- bash
# kafka-topics --create --topic test --partitions 10 --replication-factor 3 --bootstrap-server kafka-0.kafka-headless.kafka.svc.cluster.local:9092
# kafka-topics --describe --topic test --bootstrap-server kafka-0.kafka-headless.kafka.svc.cluster.local:9092
# kafka-producer-perf-test --topic test --num-records 10000 --throughput 10000 --record-size 100 --producer-props bootstrap.servers=kafka-0.kafka-headless.kafka.svc.cluster.local:9092
# kafka-consumer-groups --bootstrap-server kafka-0.kafka-headless.kafka.svc.cluster.local:9092 --describe --group my-group
export KO_DOCKER_REPO=ttl.sh
ko apply -f deployment.yaml
k9s
k apply -f scaledObject.yaml
kind delete cluster
```
