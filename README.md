# go-heap



```
go test . -run=. -bench=. -benchtime=5s -count 2 -benchmem
go test . -run=. -bench=. -benchtime=5s -count 2 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out

go tool pprof -http :8080 cpu.out
go tool pprof -http :8081 mem.out
go tool trace trace.out

go tool pprof $FILENAME.test cpu.out
# (pprof) list <func name>
```
