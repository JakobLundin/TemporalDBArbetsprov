FROM golang
ADD . /lieroarbetsprov/
ADD . /lieroarbetsprov/worker
ADD . /lieroarbetsprov/app
COPY ./app /lieroarbetsprov/app
WORKDIR /lieroarbetsprov/worker
RUN go get go.temporal.io/sdk
RUN go get
RUN go install
RUN go run worker/worker.go
