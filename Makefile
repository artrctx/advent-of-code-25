run:
	@go run main.go
start:
	mkdir internal/d${date} && \
	touch internal/d${date}/d${date}.go && \
	touch internal/d${date}/input.txt && \
	touch internal/d${date}/example.txt