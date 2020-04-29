FROM alpine

RUN apk add ca-certificates

RUN mkdir config

COPY config/env.json config

COPY Visitors-Management-System /

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["/Visitors-Management-System"]