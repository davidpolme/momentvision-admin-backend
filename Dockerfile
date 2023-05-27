# Establecer la imagen base
FROM golang:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias del módulo Go
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación
RUN go build -o app

# Establecer el comando por defecto para ejecutar la aplicación
CMD ["./app"]