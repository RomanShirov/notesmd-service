FROM golang:1.18-alpine

COPY . ./app
WORKDIR ./app
RUN go build -o build/app cmd/app/app.go
RUN cp .env build/
RUN mkdir build/assets && cp -r internal/web/notesmd-app/frontend/dist/. build/assets

WORKDIR ./build/
CMD [ "./app" ]

EXPOSE 8080:8080