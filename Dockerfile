# Étape de construction
FROM golang:1.21.1-alpine AS builder

# Répertoire de travail pour copier le code source
WORKDIR /app

# Copier les fichiers go.mod et go.sum
COPY go.mod go.sum ./

# Télécharger toutes les dépendances
RUN go mod download

# Copier le code source dans le conteneur
COPY . .

# Compiler l'application
RUN go build -o /app/main .

# Étape finale
FROM alpine:latest

# Répertoire de travail dans lequel le binaire sera copié
WORKDIR /root/

# Copier le binaire de l'étape de construction
COPY --from=builder /app/main .

# Exécuter le binaire
CMD ["./main"]
