FROM golang

ENV GOBIN $GOPATH/bin
# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/mux
RUN go get 	github.com/rafaSalgado/livingcost/config
RUN go get 	github.com/rafaSalgado/livingcost/dao
RUN go get 	github.com/rafaSalgado/livingcost/models


# copy the local package files to the container workspace
ADD . /go/src/github.com/rafaSalgado/livingcost

# Build the livingcost command inside the container.
#RUN go install github.com/antonioLibre/livingcost
RUN go install go/src/github.com/rafaSalgado/livingcost

# Setting up working directory
WORKDIR /go/src/github.com/rafaSalgado/livingcost


# Run microservices when the container starts.
ENTRYPOINT /go/bin/livingcost  --port 3000 --host 0.0.0.0


EXPOSE 3000
