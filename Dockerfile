FROM arm32v6/alpine:3.6
WORKDIR /app
# Now just add the binary
COPY pdf2HTML /app/
ENTRYPOINT ["/app/pdf2HTML"]
EXPOSE 3000