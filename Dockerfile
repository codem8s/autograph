FROM alpine:3.5

RUN apk --no-cache add ca-certificates bash && update-ca-certificates

COPY ./autograph /autograph
RUN chmod +x /autograph

CMD ["/autograph", "run"]