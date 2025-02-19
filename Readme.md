# Project Name

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
![Version](https://img.shields.io/badge/version-1.0.0-orange)

## 📌 Description

A command-line tool written in Go, designed to be executed either directly via CLI or as a compiled Go binary.

The goal of this small program is to perform a load test on your APIs. Simply provide it with a spec.api file containing your routes, and you're good to go! 🚀

## 🚀 Features

✅ Unleash concurrency!
✅ Describe your API in a simple way
✅ Run on any platform (I guess)

## 🏗️ Installation

### Prerequisites
Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.18 or higher)

### Steps

```sh
# Clone the repository
git clone https://github.com/your-repo/project.git
cd project

# Build the project
# Linux
env GOOS=linux go build -o run-run .
# Mac
env GOOS=darwin GOARCH=amd64 go build -o run-run-mac .
# Run the project
./run-run --help # Linux
./run-run-mac # mac
```

## 📖 Usage


```sh
# Example: Start 10 requests with unlimited concurrency using my file spec2.api
./run-run -file spec2.api -r 10

# Example: Start 10 requests, 1 request per second
./run-run -file spec2.api -r 10 -freq "s" -rate 1

# Example: Start 10 requests with a maximum concurrency of 2
./run-run -file spec2.api -r 10 -c 2
```

## 🛠️ Configuration
###  spec.api exemple
{
  "routes": [
    {
      "method": "GET",
      "url": "http://localhost:8000/load/get",
      "headers": {
        "Authorization": "Bearer <token>",
        "Prout": "Prout"
      },
      "body": null,
      "description": "Retrieve a list of users"
    },
        {
      "method": "POST",
      "url": "http://localhost:8000/load/post",
      "headers": {
        "Authorization": "Bearer <token>",
        "Prout": "Prout"
      },
      "body": {
        "test":"flute"
      },
      "description": "Retrieve a list of users"
    },
        {
      "method": "PUT",
      "url": "http://localhost:8000/load/put",
      "headers": {
        "Authorization": "Bearer <token>",
        "Prout": "Prout"
      },
      "body": {
        "name": "raph",
        "email": "r@gmail.com"
      },
      "description": "Retrieve a list of users"
    }
  ]
}

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

## 👤 Authors

- **Raphael Obadia** 