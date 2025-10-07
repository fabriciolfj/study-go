FROM golang:1.24 AS build

RUN useradd --create-home --shell /bin/bash appbuild

WORKDIR /home/appbuild

# Copia arquivos Go
COPY go.mod go.sum ./

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN go build -o /home/appbuild/bin/app ./main.go

FROM ubuntu:24.04

RUN useradd --create-home --shell /bin/bash apprun

# Copia o binário compilado
COPY --from=build /home/appbuild/bin/app /home/apprun/app

# Define o usuário
USER apprun
WORKDIR /home/apprun

CMD ["./app"]