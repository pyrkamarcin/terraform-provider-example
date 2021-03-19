default: clean build install

clean:
	rm -rf .terraform
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate
	rm -f  ~/.terraform.d/plugins/github.com/pyrkamarcin/example/0.1/darwin_amd64/terraform_provider_example/terraform-provider-example

build:
	go build -o terraform-provider-example

install:
	mv terraform-provider-example ~/.terraform.d/plugins/github.com/pyrkamarcin/example/0.1/darwin_amd64/

test: clean build install
	terraform init
	TF_LOG=DEBUG terraform plan
	TF_LOG=DEBUG terraform apply -auto-approve
