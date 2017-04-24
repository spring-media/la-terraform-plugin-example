# la-terraform-plugin-example
Terraform example plugin

### Build 

    go build -o terraform-provider-dummy

### Use
Install the plugin https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin

    providers {
        dummy = "$GOPATH/src/github.com/WeltN24/la-terraform-plugin-example/terraform-provider-dummy"
    }


or copy it into the project directory

### For random joke
    
    data "chucks_jokes" "foo" {}
    output "joke" {
        value = "${data.dummy_server.foo.joke}"
    }

### For joke with id

    data "chucks_jokes" "foo" {
        id = 103
    }

    output "joke" {
        value = "${data.dummy_server.foo.joke}"
    }
