test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... && go test -coverprofile coverage.out -cover ./... # use -v for verbose
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.out

bench:
	#go test -bench=. -benchtime 100x
	go test -bench=. -benchmem
	#go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out