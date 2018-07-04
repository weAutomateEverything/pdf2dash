FROM phusion/baseimage
WORKDIR /app
COPY pdf2dash /app/
ENTRYPOINT ["/app/pdf2dash"]
EXPOSE 3000