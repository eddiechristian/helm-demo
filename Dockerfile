FROM helm_demo_build:latest
COPY go/bin/main /usr/local/bin
ENTRYPOINT ["/usr/local/bin/main"]


