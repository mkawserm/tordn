# tordn 
-------------------------------------------------

[![Build Status](https://travis-ci.com/mkawserm/tordn.svg?branch=master)](https://travis-ci.com/mkawserm/tordn)
[![Go Report Card](https://goreportcard.com/badge/github.com/mkawserm/tordn)](https://goreportcard.com/report/github.com/mkawserm/tordn) 
[![Coverage Status](https://coveralls.io/repos/github/mkawserm/tordn/badge.svg?branch=master)](https://coveralls.io/github/mkawserm/tordn?branch=master)
-------------------------------------------------

# The Onion Router Domain Name

Package tordn provides the ability to generate tor version 3 onion address for the hidden service

# Example

```
package main

import (
	"fmt"
	"github.com/mkawserm/tordn"
)

func main() {

	v3domainName := &tordn.V3{}

	// generate random v3 tor domain name
	publicKey, privateKey, onionAddress, err := v3domainName.GenerateTORDomainName(nil)
	if err == nil {
		fmt.Printf("Public Key:")
		fmt.Println(publicKey)

		fmt.Printf("Private Key:")
		fmt.Println(privateKey)

		fmt.Printf("Onion Address:")
		fmt.Println(string(onionAddress))
	}
}

```

# DEVELOPER NOTE 
    1. Benchmark memory profile
        
        Create memory benchmark
         
        `go test -bench=. -memprofile=mem.out`
    
        Top  profile
             
        `go tool pprof -alloc_objects -top -cum mem.out`
        
        Line by line profile
        
        `go tool pprof -alloc_objects -list=<BenchmarkName> mem.out`
    
    2. Escape Analysis
        
        `go test -gcflags '-m -m'`
        
    3. Test coverage
    
        `go test -coverprofile=cover.out`
        
    4. Race condition detection
    
        `go test -race`

