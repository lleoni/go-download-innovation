FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-download-innovation"]
COPY ./bin/ /