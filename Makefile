prepare:
	kubectl apply -f hack/cluster-role.yaml
	kubectl apply -f hack/deploy-dev.yaml

check-prepare:
	kubectl get pod -n mount-log

build:
	go build -o loghacker

run: build
	kubectl cp -n mount-log loghacker  mount-dev:/tmp
	kubectl exec -it -n mount-log mount-dev -- bash

run-check: build
	kubectl cp -n mount-log loghacker  mount-var-log:/tmp
	kubectl exec -it -n mount-log mount-var-log -- bash

docker:
	docker build . -t shadowstar884/loghacker

docker-push: docker
	docker push shadowstar884/loghacker

deploy:
	kubectl apply -f deploy-loghacker.yaml