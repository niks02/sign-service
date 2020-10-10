# Signature Service

This repository contains the Signature Service. 

## I just cloned the repository. What next?

### Building Code
From the repo base directory, run the following command
```bash
# To build
go build .
# To run the server
./signature_service
```

## Repository structure
* `README.md` - the first file you have to read for this repository
* `go.mod` golang's module file that stores exact version of the dependencies. It is often accompanied by `go.sum` which stores the exact version of the dependencies'.
* `main.go` - the entry point to our service.
* `config/` directory stores the business logic for the config API. `config.json` is present in this directory where port and keypair is present. 
* `signature/` directory stores the business logic for the signature API.
* `transaction/` directory stores the business logic for the transaction API.
* `Assignment_Part2/` directory contains the answers to Part2 questions
**Note that** unit tests are not included due to time constraints

## Code Tested on 
* macOS Mojave Version 10.14.6
* Go Version 1.14
