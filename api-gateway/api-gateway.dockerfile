FROM alpine:latest
RUN mkdir /app
COPY apiGatewayApp /app
CMD ["/app/apiGatewayApp"]