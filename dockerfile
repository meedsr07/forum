FROM golang:1.22.3-alpine

# Metadata
LABEL source="https://github.com/meedsr07/forum.git"
LABEL maintainer="ouaaitalla"
LABEL project="Forum"
LABEL version="1.0"
LABEL description="This image is for forum project"
LABEL created="16-04-2026"

#set working directory
WORKDIR /forum

#copy project diractory from host to image
COPY . . 

#compile go code
RUN go build -o app

#Expose web server port
EXPOSE 8088

#default command 
CMD ["./forum"]