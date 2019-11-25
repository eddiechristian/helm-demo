build:
	docker run --rm -v /Users/eddie.christian/source/helm-demo/:/root helm_demo_build:latest /root/build.sh

docker:
	docker build -t helm_demo_go .

test:
	docker run --rm -p 8080:9990 helm_demo_go:latest 

install-helm-demo:
	kubectl config use-context docker-for-desktop
	helm install -n helm-demo --namespace helm-demo --wait ./charts/helm_demo/

delete-helm_demo:
	helm del --purge helm-demo
	kubectl delete pvc data-helm-demo-postgresql-0 -n helm-demo

deploy-helm-demo-db:
	kubectl cp db.sql default/helm-demo-postgresql-0:/tmp/db.sql
	echo "#!/bin/bash" > db.sh
	echo 'sql=$$(cat /tmp/db.sql)' >> db.sh
	echo "export PGPASSWORD=$$(kubectl get secret helm-demo-postgresql -o json | jq -r '.data."postgresql-password"' | base64 --decode)" >> db.sh
	echo 'psql -U postgres -c "CREATE DATABASE helm_demo;"' >> db.sh
	echo 'psql -U postgres -d helm_demo -c "$$sql"' >> db.sh
	chmod +x db.sh
	kubectl cp db.sh default/helm-demo-postgresql-0:/tmp/db.sh
	kubectl exec helm-demo-postgresql-0  /tmp/db.sh
	