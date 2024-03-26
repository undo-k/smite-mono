FROM alpine:latest
RUN mkdir /app
COPY mockApiApp /app
CMD ["/app/mockApiApp"]