# Building

    docker build -t gofizzbuzz .

# Starting

    docker run -p 9000:9000 -i -t gofizzbuzz

# Developping

This application is a revel app : you'll need revel checked out in your GOPATH :

    go get -u github.com/revel/revel
    go get -u github.com/revel/cmd/revel

To start the application in dev mode :

    revel run -a github.com/tcurtil/go-fizz-buzz


