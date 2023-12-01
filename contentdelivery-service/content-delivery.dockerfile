#Building stage
FROM golang:1.21-alpine as builder

#run the command on the docker image we are building
RUN mkdir /app

COPY . /app 

WORKDIR /app 

#Build go code 
RUN CGO_ENABLE=0 go build -o contentApp ./cmd

#run the chmod command and add the executable flag
RUN chmod +x /app/contentApp 


##Running Stage
FROM alpine:latest

RUN mkdir /app 

#copy files from the builfer stage to /app 
COPY --from=builder /app/contentApp /app 

CMD [ "/app/contentApp" ]