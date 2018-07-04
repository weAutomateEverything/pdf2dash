FROM phusion/baseimage
WORKDIR /app
COPY staticPages /app/
COPY pdfFile.pdf /app/
ENTRYPOINT ["/app/pdf2dash"]
EXPOSE 3000