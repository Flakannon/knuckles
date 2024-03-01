GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/battlelambda cmd/battlelambda/main.go

docker-compose down

docker-compose up -d --build

cd ./terraform

tflocal init
tflocal plan
tflocal apply -auto-approve
