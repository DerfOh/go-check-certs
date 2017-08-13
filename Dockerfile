# Requirements for building:
#   Run ./dockerbuild.sh
#   This will cross compile the go binary to linux then
#   runs docker build to add the binary

FROM alpine:latest

MAINTAINER DerfOh <fredrick.p@outlook.com>

# Update packages and ca-certificates
RUN apk update
RUN apk add ca-certificates

# Set up results directory and files

RUN mkdir -p /app/results

### Set env variables ##
ENV CHECK_SIG_ALG true

ENV CONCURRENCY 8

ENV DAYS 0

ENV MONTHS 0

ENV YEARS 1

ENV OUTPUT false

ENV SERVE false

# Results path must be absolute path
# ENV RESULTS /app/results/

EXPOSE 8080 8080

## Move files over to container ##
COPY go-check-certs /app

COPY index.html /app

COPY hosts.txt /app

ENTRYPOINT ["sh", "-c", "/app/go-check-certs -check-sig-alg=${CHECK_SIG_ALG} -concurrency=${CONCURRENCY} -hosts=/app/hosts.txt -days=${DAYS} -months=${MONTHS} -years=${YEARS} -serve=${SERVE} -output=${OUTPUT} -results=${RESULTS}"]
# Example Usage: docker run -e SERVE='true' -e YEARS=1 -it -p 8080:8080 derfoh/go-check-certs