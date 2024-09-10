# Utiliser l'image officielle de Go
FROM golang:1.23.1

# Définir le répertoire de travail dans le conteneur
WORKDIR /social-network

# Installer SQLite3
RUN apt-get update && \
    apt-get install -y sqlite3 libsqlite3-dev

# Copier les fichiers go.mod et go.sum dans le répertoire de travail
COPY ./go.mod .

# Télécharger les dépendances Go
RUN go mod download

# Copier tout le code source dans le conteneur
COPY . .

# Exposer le port (si nécessaire)
EXPOSE 8080

# Empêcher le conteneur de se terminer immédiatement en démarrant un shell interactif
CMD ["tail", "-f", "/dev/null"]