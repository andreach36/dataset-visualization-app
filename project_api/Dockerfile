# Etapa de construcción
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar go.mod y go.sum
COPY ./go.mod ./go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Etapa de ejecución
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Crear el directorio de trabajo
WORKDIR /app

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/main .

# Exponer el puerto
EXPOSE 8080

# Establecer las variables de entorno
ENV PORT=${PORT}
ENV GIN_MODE=${GIN_MODE}

# Ejecutar el binario
CMD ["./main"]
