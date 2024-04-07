FROM golang:1.21-bullseye

RUN wget -O gf "https://github.com/gogf/gf/releases/latest/download/gf_$(go env GOOS)_$(go env GOARCH)" && chmod +x gf && ./gf install -y && rm ./gf
