FROM golang:1.24-bullseye

ENV PATH=/usr/local/bin:$PATH

ENV LANG=C.UTF-8

# Set the Current Working Directory inside the container
WORKDIR /app

RUN mkdir /app/temp
RUN cd /app/temp

# Copy the source from the current directory to the Working Directory inside the container
COPY . .


# run gofmt on all source files
RUN gofmt -s -w .

# go mod download
RUN go mod download

RUN ls -la

# Build the Go app
RUN go build -o main cmd/App/main.go
# Expose port 8080 to the outside world

EXPOSE 8080


RUN cd /app
RUN rm -rf /app/temp

ENTRYPOINT ["./main"]