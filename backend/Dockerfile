# Use uma imagem base do Go
FROM golang:1.21.4

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie o go.mod e go.sum e baixe as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte para o contêiner
COPY . .

# Compile o binário da aplicação
RUN go build -o server cmd/server/main.go

# Cria o diretório de uploads
RUN mkdir -p /app/uploads

# Exponha a porta na qual a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./server"]
