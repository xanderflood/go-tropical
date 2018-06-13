commands := $(shell ls cmd/ | awk '{split($$0,a,"/"); print a[1]}' | tr '\n' ' ')

all: $(commands)

$(commands):
	@echo "Building command $@"
	go build cmd/$@/main.go
	mv main bin/$@
