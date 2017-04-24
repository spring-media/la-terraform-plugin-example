# la-terraform-plugin-example
Terraform example plugin

### Build 

    go get github.com/WeltN24/la-terraform-plugin-example

### Use
Install the plugin https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin

Example ~/.terraformrc: 

    providers {
        chucks = "$GOPATH/bin/la-terraform-plugin-example"
    }


or copy it into the project directory

### For random joke
    
    data "chucks_jokes" "foo" {}
    output "joke" {
        value = "${data.chucks_jokes.foo.joke}"
    }

### For joke with id

    data "chucks_jokes" "foo" {
        id = 103
    }

    output "joke" {
        value = "${data.chucks_jokes.foo.joke}"
    }
