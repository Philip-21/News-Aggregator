#Building stage
FROM golang:1.21-alpine as builder

#run the command on the docker image we are building
RUN mkdir /app

COPY . /app 

WORKDIR /app 

#Build go code 
RUN CGO_ENABLE=0 go build -o newsApp ./cmd

#run the chmod command and add the executable flag
RUN chmod +x /app/newsApp 


##Running Stage
FROM alpine:latest

RUN mkdir /app 

#copy files from the builfer stage to /app 
COPY --from=builder /app/newsApp /app 

EXPOSE 8001

CMD [ "/app/newsApp" ]