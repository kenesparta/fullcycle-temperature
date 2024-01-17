FROM golang:1.21.5-bookworm as builder
ARG ENV_ARG
ENV ENVIRONMENT=$ENV_ARG
WORKDIR /app
COPY . .
RUN make init && make build

FROM scratch
ARG ENV_ARG
ENV ENVIRONMENT=$ENV_ARG
COPY --from=builder /app/server .
COPY --from=builder /app/"env.$ENVIRONMENT.json" .
EXPOSE 8080
CMD ["./server"]
