FROM golang
WORKDIR /opt/ascii-art
COPY . .
RUN go mod download
EXPOSE 8080
RUN go build -o ascii
CMD [ "./ascii" ]