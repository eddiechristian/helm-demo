build:
	docker run --rm -v /Users/eddie.christian/source/helm-demo/:/build helm_demo_build:v1 /build/build.sh

docker:
	docker build -t helm_demo_go .

install-helm-demo:
	kubectl config use-context docker-for-desktop
	helm install -n helm-demo --namespace helm-demo --wait ./charts/helm_demo/
	helm get helm-demo
test:
	yq w --inplace charts/helm_demo/templates/configmap.yaml data.pghost $(shell kubectl get svc -n helm-demo helm-demo-postgresql -o json | jq -r '.spec.clusterIP')
	yq w --inplace charts/helm_demo/templates/deployment.yaml spec.template.metadata.annotations.configHash $(shell cat charts/helm_demo/templates/configmap.yaml | md5)
delete-helm-demo:
	helm del --purge helm-demo
	kubectl delete pvc data-helm-demo-postgresql-0 -n helm-demo

configure-metallb:
	kubectl get po -l app=helm-demo-pod -n helm-demo -o json | jq '.items[] | { podip: .status.podIP}'

deploy-helm-demo-db:
	kubectl cp db.sql helm-demo/helm-demo-postgresql-0:/tmp/db.sql
	echo "#!/bin/bash" > db.sh
	echo 'sql=$$(cat /tmp/db.sql)' >> db.sh
	echo "export PGPASSWORD=$$(kubectl get secret helm-demo-postgresql -n helm-demo -o json | jq -r '.data."postgresql-password"' | base64 --decode)" >> db.sh
	echo 'psql -U postgres -c "CREATE DATABASE helm_demo;"' >> db.sh
	echo 'psql -U postgres -d helm_demo -c "$$sql"' >> db.sh
	chmod +x db.sh
	kubectl cp db.sh helm-demo/helm-demo-postgresql-0:/tmp/db.sh
	kubectl exec -n helm-demo helm-demo-postgresql-0  /tmp/db.sh
	