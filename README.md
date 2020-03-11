# tordn
-------------------------------------------------

The Onion Router Domain Name


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

