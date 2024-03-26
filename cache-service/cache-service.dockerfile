FROM alpine:latest
RUN mkdir /app
COPY cacheApp /app
CMD ["/app/cacheApp"]