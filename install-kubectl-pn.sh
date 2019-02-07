#/bin/bash

go build cmd/kubectl-pn.go
sudo chmod +x ./kubectl-pn
sudo mv ./kubectl-pn /usr/local/bin