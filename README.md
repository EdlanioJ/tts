<h1 align="center">Text To Speech with GRPC Server</h1>

A [GO](https://go.dev/) application to convert written text into spoken voice and download it in mp3 file. This project is made with Google speech synthesizes.</p>

## Getting Started

Require [Docker](https://www.docker.com/) to run.
Run the development server:

```bash
# create .env file with .env.example data

# start docker container 

docker-compose up --build
docker exec -it tts bash  

# build client and server for linux

make build.server
make build.client

# run test

make test
```

Start on Linux:

```bash
# start server

./tts-server

# Get audio

./say -text="Hello, World!" -lang="en"

# Get audio from file

./say -file="test.txt" lang="en"

# Set output

./say -text="Hello, World!" -lang="en" -out="output.mp3"

# Get Help

./say -help 
```
