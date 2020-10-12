package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAwsHelloWorld(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		// 1 - The path to where our Terraform code is located
		TerraformDir: "../modules/hello-world-aws",
	}

	// 5 - At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// 2 - Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// 3 - Run `terraform output` to get the IP of the instance
	publicIP := terraform.Output(t, terraformOptions, "public_ip")

	// 4 - Make an HTTP request to the instance and make sure we get back a 200 OK with the body "Hello, World!"
	url := fmt.Sprintf("http://%s:8080", publicIP)
	http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello, World!", 30, 5*time.Second)
}
