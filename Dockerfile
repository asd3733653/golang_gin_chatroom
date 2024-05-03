
# 使用官方的 Golang 映像作為基礎映像
FROM golang:1.22.0-alpine

# 在容器中建立一個目錄來存儲我們的應用程式
RUN mkdir /app

# 複製應用程式的程式碼到容器中的 /app 目錄
ADD . /app

# 設置當前的工作目錄為 /app 目錄
WORKDIR /app

COPY .env .env

# 下載並安裝依賴
RUN go mod download

# 編譯 Go 應用程式
RUN go build -o main .

# 告訴 Docker 執行該容器提供服務的端口號
EXPOSE 8080

ENV GIN_PORT=32161

# 設置容器啟動後執行的命令
CMD ["/app/main"]
