# Running test cases
Step1: go to dir "gltest" in terminal with "cd"

Setp2: run testing cmd

	like "$go test -v glclient_test.go" running test cases in "glclient_test.go",
	and "$go test -v" running test cases in "*.go" of current dir,
	and "-v" means output details of testing cases' executing process

Step3:review the testing result
```
run current dir's testing code: $ go test 
run current dir's testing code and show process: $ go test -v
run testing code in target go file: $ go test xxx.go 
run testing code in target go file and show process: $ go test -v xxx.go 
run test of standard libs: $ go test std
get test cases' coverage: $ go test -cover
run Parallel Test: $ go test -bench=".*"
```
1.Parallel Test
```
Run: $ go test -bench=".*"
Result:
	PASS
	#         function name	       CPU  loop times	 each loop's time
	Benchmark_GETReqWithURLParallel-4	  200000	      9054 ns/op
	Benchmark_GETReqWithURL-4        	 1000000	      2050 ns/op
	ok  	github.com/coder4869/golibs/gltest	14.687s
```

# MORE
```
GO Testing Reference:http://studygolang.com/articles/1889
or running "$ go test -help" in terminal
```