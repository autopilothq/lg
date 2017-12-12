
bench:
	go test -run=none -bench=. -benchmem

benchOut:
	go test -run=none -bench=. -benchmem > fancyBench.txt

prof:
	go test -run=none -bench=BenchmarkFancy -cpuprofile=cprof -memprofile=mem.out -benchmem
	go tool pprof --text benchmark.test cprof

compare: benchOut
	benchcmp reflectionBench.txt fancyBench.txt

compareViz: benchOut
	benchcmp reflectionBench.txt fancyBench.txt | benchviz > bench.svg

installTools:
	go get -u github.com/ajstarks/svgo/benchviz \
				golang.org/x/tools/cmd/benchcmp
