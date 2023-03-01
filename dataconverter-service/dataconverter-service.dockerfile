FROM alpine:latest

RUN mkdir /app

COPY dataConverterApp /app

CMD [ "./app/dataConverterApp" ]