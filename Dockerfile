FROM helm_demo_build:latest
COPY go/bin/main /usr/local/bin
RUN apt-get install -y postgresql-client
ENTRYPOINT ["/usr/local/bin/main"]


