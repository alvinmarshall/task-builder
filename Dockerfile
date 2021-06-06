FROM golang:1.13.8-stretch as builder
WORKDIR /taskbuilder
COPY . /taskbuilder

ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_NAME=task

ENV PORT=8888
ENV JWT_SECRET=secret
ENV JWT_EXPIRES=60
ENV APPLICATION_NAME=task-builder

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o taskbuilder cmd/task.go

FROM gcr.io/distroless/static
WORKDIR /cip
COPY --from=builder /taskbuilder .

ENTRYPOINT ["./taskbuilder"]