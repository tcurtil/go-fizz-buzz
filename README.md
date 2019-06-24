# Building

    docker build -t gofizzbuzz .

# Starting

    docker run -p 9000:9000 -i -t gofizzbuzz

# Solving the fizzbuzz exercise :

## Subject:

Your goal is to implement a web server that will expose a REST API endpoint that: 
Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
Ready for production
Easy to maintain by other developers

Bonus Question :
- Add a statistics endpoint allowing users to know what the most frequent request has been. 
This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

## info on the solution
- first endpoint : /fizzbuzz?limit=[limit]&int1=[int1]&int2=[int2]&str1=[str1]&str2=[str2]
- second endpoint (stat) : /stats

# Developping

This application is a revel app : you'll need revel checked out in your GOPATH :

    go get -u github.com/revel/revel
    go get -u github.com/revel/cmd/revel

To start the application in dev mode :

    revel run -a github.com/tcurtil/go-fizz-buzz


