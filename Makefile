bin=silverbullet

build:
	go build -o ${bin} cmd/${bin}/*go
check:
	go fmt cmd/${bin}/*go
