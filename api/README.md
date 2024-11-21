# ResumeKraft API Server

API server is serving the API for ResumeKraft

## Usage
Check the .env.example copy the file and rename it to .env
```
mv .env.example .env
```

Then, install the dependencies
```
go mod tidy
```

Run the migrations
```
go run main.go migrate up
```

Run the server
```
go run main.go api
```