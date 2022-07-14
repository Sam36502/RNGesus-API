FROM ubuntu:18.04

EXPOSE 777

ENV HOLY_LAND_KEY=0a63fd62d9053d16b8b7d6ebd140ac55

RUN mkdir /app
COPY dist/rngesus /app/rngesus
WORKDIR /app

RUN apt update
RUN apt install -y ca-certificates

ENTRYPOINT [ "/app/rngesus" ]