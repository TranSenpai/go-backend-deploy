FROM docker.io/library/golang:1.26.4-alpine AS build

# Tạo folder để chứa source vào build, bắt đầu từ WORKDIR này trờ xuống sẽ làm việc trong folder đó
WORKDIR /app

# Coppy go.mod go.sum ./ để  Docker chụp danh sách thư viện lại trước để check xem có thay đổi không ? Nếu có mới dowload không thì dùng CACHED. Nếu không COPY go.mod go.sum ./ mà COPY . . xong RUN go mod download thì sửa code mà không tải thêm thư viện Docker cũng download lại từ đầu mà không dùng cached
COPY go.mod go.sum ./
RUN go mod download

# Copy hết các file code vào để  build, dáu . đầu là adrress gốc, dấu . sau là address các file copy 
COPY . .

# Lệnh chạy trong lúc buildtime (Run build)
RUN go build -o go-backend ./cmd

# Lệnh chạy lúc runtime
# CMD ["./go-backend"]

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/go-backend .

CMD ["./go-backend"]



# 2.3 GB

# dockerignore: 2.18 GB

# stage: 69.86 MB