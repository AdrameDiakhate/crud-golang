 #Utilisation de golang comme image de base
# Le GOPATH par défaut de cette image est /go.
FROM golang:latest
WORKDIR /app
# Copie des sources de notre projet dans le container,
# dans notre cas le main.go
COPY go.mod .
COPY go.sum . 

# Lancement de la compilation avec go install
RUN go mod download 
COPY . . 
EXPOSE 8080
RUN go build
CMD [ "pratice-crud-go-main/pratice-crud-go-main" ]

# Définissions du point d'entré de notre programme compilé
#ENTRYPOINT /go/bin/hello
 

# Le port sur lequel notre serveur écoute
#EXPOSE 8080