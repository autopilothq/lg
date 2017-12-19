GOPATH=$(shell go env GOPATH)
BENCHCMP=${GOPATH}/src/golang.org/x/tools/cmd/benchcmp
BENCHVIS=${GOPATH}/src/github.com/ajstarks/svgo/benchviz

compare: bench.out
	benchcmp oldBench.out bench.out

compareviz: bench.svg

bench.svg: bench.out
	benchcmp oldBench.out bench.out | benchviz > bench.svg

bench: installTools
	go test -run=NONE -bench=. -benchmem

bench.out: installTools
	go test -run=NONE -bench=. -benchmem > bench.out

prof: installTools
	go test -run=NONE -bench=BenchmarkPlainTextFieldLogging -cpuprofile=cpuprof.out -memprofile=memprof.out -benchmem

cpuprof.out: prof
	go tool pprof --text benchmark.test cpuprof.out

cpuprof: cpuprof.out

memprof.out: prof
	go tool pprof --text benchmark.test memprof.out

memprof: memprof.out

installTools: ${BENCHCMP}/.sentinel ${BENCHVIS}/.sentinel

.PRECIOUS: ${BENCHCMP}/.sentinel # otherwise make (annoyingly) deletes it
${BENCHCMP}/.sentinel:
		 go get -u golang.org/x/tools/cmd/benchcmp
		 touch $@

.PRECIOUS: ${BENCHVIS}/.sentinel # otherwise make (annoyingly) deletes it
${BENCHVIS}/.sentinel:
		go get -u github.com/ajstarks/svgo/benchviz
		touch $@

clean:
	rm -f bench.out bench.svg cpuprof.out memprof.out ${BENCHCMP}/.sentinel ${BENCHVIS}/.sentinel

.PHONY: clean compare compareviz bench prof cpuprof memprof
