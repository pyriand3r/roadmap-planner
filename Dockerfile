FROM golang:1.22-bookworm AS gobuilder
RUN apt update && \
    apt install -y git nodejs npm && \
    apt clean

WORKDIR /root
COPY . .

RUN make build

FROM debian:bookworm
COPY --from=gobuilder /root/build ./rplanner
VOLUME /rplanner/data
WORKDIR /rplanner

CMD ["./rplanner"]
