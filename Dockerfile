FROM golang

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/gorilla/mux
RUN go get github.com/RafaSalgado/livingcost/config
RUN go get github.com/RafaSalgado/livingcost/models


# copy the local package files to the container workspace
ADD . /go/src/github.com/RafaSalgado/livingcost

# Build the livingcost command inside the container.
#RUN go install github.com/antonioLibre/livingcost
RUN go install github.com/RafaSalgado/livingcost

# Setting up working directory
WORKDIR /go/src/github.com/RafaSalgado/livingcost


# Run microservices when the container starts.
ENTRYPOINT /go/bin/livingcost


EXPOSE 4001
