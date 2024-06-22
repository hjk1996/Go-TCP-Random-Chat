# Build stage
FROM golang:1.21.5 as builder
WORKDIR /app
# 모듈 다운로드와 캐시 최적화를 위해 go.mod와 go.sum 파일을 먼저 복사합니다.
COPY go.mod go.sum ./
RUN go mod download
# 나머지 소스 파일을 복사하고 빌드합니다.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o myapp cmd/server/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
# 빌드된 바이너리만 복사합니다.
COPY --from=builder /app/myapp .
# 실행 권한은 빌드 단계에서 설정되므로 필요하지 않습니다.
# RUN chmod +x myapp 
# 디버깅을 위해 ls 명령어를 추가하여 파일 존재 여부를 확인합니다.
RUN ls -la /app
CMD ["/app/myapp"]


