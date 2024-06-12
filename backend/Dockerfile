# Golang resmi imajını kullan
FROM golang:1.18-alpine

# Çalışma dizinini ayarla
WORKDIR /app

# Go mod ve sum dosyalarını çalışma dizinine kopyala
COPY go.mod ./
COPY go.sum ./

# Bağımlılıkları indir
RUN go mod download

# Uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle
EXPOSE 3000


