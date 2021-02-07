FROM golang
ADD . /lieroarbetsprov/webservice
WORKDIR /lieroarbetsprov/webservice
RUN go get go.temporal.io/sdk
RUN go get
RUN go install
RUN go run http.go
