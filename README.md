# docker-events
Send the events docker generates to NATS.

There are a few requirements that are put on you before you can use this piece of software.

1. You need to build it. ```go build``` will fix that.
2. You needs to have access to nats. (http://nats.io/download/)
2. The binary needs you to pipe docker events to it AND the events has to be formated as json. Like this ```docker events --format '{{json .}}' | binary -nats=<nats url>```.

And now magic will happen.

## Docker events crash course

Docker events have a Type and an Action. These are combined to build the "routing key" (in traditional pubsub)

The ```Type``` will be something like ```container``` or ```network```. So one can say container commands will generate container events. But it does not end there, they will also generate a network event, if the container are added to a network (which most are).

The ```Action``` might be something like ```start``` and kind of reflect the docker command that was ran. But in some cases a single command generates a handful of events.

With this brief intro, an event of ```Type=container``` and ```Action=start``` will be put on topic ```container.start```. The event will be what ever docker gave us in the first case. Expect quite alot of data. It's all yours.