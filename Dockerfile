FROM alpine:3.6
WORKDIR /app
COPY pdf2dash /app/
ENTRYPOINT ["/app/pdf2dash"]
EXPOSE 3000