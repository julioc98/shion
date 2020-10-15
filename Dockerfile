FROM golang
WORKDIR /shion
COPY . .
CMD [ "make","run/api"]