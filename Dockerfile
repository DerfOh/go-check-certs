# Requirements for building:
#   Run ./dockerbuild.sh
#   This will cross compile the go binary to linux then
#   runs docker build to add the binary

# FROM alpine:latest
FROM ubuntu:latest

MAINTAINER DerfOh <fredrick.p@outlook.com>

# Set env variables
ENV CHECK_SIG_ALG true

ENV CONCURRENCY 8

# TODO: Figure out best way to handle hosts file
# ENV HOSTS

ENV DAYS 0

ENV MONTHS 0

ENV YEARS 1

ENV OUTPUT false

ENV SERVE false

EXPOSE 8080 8080

# Move files over to container

COPY go-check-certs /usr/bin

COPY index.html /usr/bin

COPY hosts.txt /usr/bin

# Set up results directory and files

RUN mkdir /usr/bin/results

RUN touch /usr/bin/results/results.csv

# Update packages and ca-certificates
# RUN apk update

# RUN apk add ca-certificates

RUN apt-get update -y
RUN apt-get install  -y ca-certificates

ENTRYPOINT ["sh", "-c", "/usr/bin/go-check-certs -check-sig-alg=${CHECK_SIG_ALG} -concurrency=${CONCURRENCY} -hosts=/usr/bin/hosts.txt -days=${DAYS} -months=${MONTHS} -years=${YEARS} -serve=${SERVE} -output=${OUTPUT} -results=${RESULTS}"]
# ENTRYPOINT ["sh"]