## Copyright 2021 Paingha Joe Alagoa. All rights reserved.
## Use of this source code is governed by a BSD-style
## license that can be found in the LICENSE file.

# select image
FROM golang:1.15-alpine
WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download
ADD . .
WORKDIR /app/server
COPY example.env .
RUN go mod vendor
RUN go build -o server
# Command to run the executable
CMD ["./server"]
