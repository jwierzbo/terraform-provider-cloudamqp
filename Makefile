NAME="terraform-provider-cloudamqp"
VERSION="1.0.0"

.PHONY: terraform-provider-cloudamqp
terraform-provider-cloudamqp:
	go install github.com/jwierzbo/${NAME}/cmd/${NAME}

.PHONY: create-release
create-release:
	mkdir -p build/${NAME}/linux-amd64
	mkdir -p build/${NAME}/darwin-amd64

	GOOS=linux GOARCH=amd64 go build -o build/${NAME}/linux-amd64/${NAME} \
		github.com/jwierzbo/${NAME}/cmd/${NAME}

	GOOS=darwin GOARCH=amd64 go build -o build/${NAME}/darwin-amd64/${NAME} \
		github.com/jwierzbo/${NAME}/cmd/${NAME}

.PHONY: install-local
install-local:
	go build -o terraform-provider-cloudamqp github.com/jwierzbo/${NAME}/cmd/${NAME}
	mkdir -p ~/.terraform.d/plugins/
	mv terraform-provider-cloudamqp ~/.terraform.d/plugins/
