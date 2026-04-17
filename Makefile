APP     := pixelcounter
OUTDIR  := dist

.PHONY: build run clean all windows linux mac

build:
	go build -o $(APP) .

run:
	go run .

windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTDIR)/$(APP)-windows-amd64.exe .

linux:
	GOOS=linux GOARCH=amd64 go build -o $(OUTDIR)/$(APP)-linux-amd64 .

mac:
	GOOS=darwin GOARCH=amd64 go build -o $(OUTDIR)/$(APP)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o $(OUTDIR)/$(APP)-darwin-arm64 .

all: windows linux mac

clean:
	rm -f $(APP)
	rm -rf $(OUTDIR)
