FROM alpine:latest
RUN apk update
RUN apk add docker
COPY docker-events /srv/events/events
WORKDIR "/srv/events"
CMD ["./events"]