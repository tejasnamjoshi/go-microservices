# Go Microservices

### This repository is for all my go micro-services practice projects

Steps to run APIs
 - Go to the desired project's directory.
 - Open a terminal
 - Run `go run main.go`
 - Open a new terminal
 - Use cURL to make API requests
    - `curl localhost:9090 | jq` - Makes a GET request on port 9090
    - `curl -X POST localhost:9090 -d '{required_data}'` - Makes a POST request on port 9090
    - `curl -X PUT localhost:9090/{id} -d '{required_data}'` - Makes a PUT request on port 9090