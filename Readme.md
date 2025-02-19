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
```
# Clone the repository
git clone git@github.com:RaphOb/load_tester_cli.git
cd load_test_cli
```

### Using the Precompiled Binaries (Recommended)
If you just want to use the application, you can directly run the precompiled binary for your platform:

- **Linux**: Use `run-run`
- **Mac**: Use `run-run-mac`

There's no need to install Go on your machine—just download the binary and start using it!

### Building from Source (For Developers)
If you want to modify the code or compile a binary for Linux, Mac or Windows, you need to install the following prerequisites:

- [Go](https://golang.org/doc/install) (version 1.18 or higher)

### Steps to Build from Source

```sh
# Build the project
# Linux
env GOOS=linux go build -o run-run .
# Mac
env GOOS=darwin GOARCH=amd64 go build -o run-run-mac .
# Run the project
./run-run --help # Linux
./run-run-mac --help # mac
```

## 📖 Usage
### Options
-c : integer. Max Concurrency.  
-file: string. File name (path)  
-freq: string. Allowed freq are ms, s min  
-rate: int. Number of rate desired  
-r: int. Number of rate desired  
### Example
```sh
# Example: Start 10 requests with unlimited concurrency using my file spec2.api
./run-run -file spec2.api -r 10

# Example: Start 10 requests, 1 request per second
./run-run -file spec2.api -r 10 -freq "s" -rate 1

# Example: Start 10 requests with a maximum concurrency of 2
./run-run -file spec2.api -r 10 -c 2
```
## 📊 Result
```
╭─────────────────────────────────┬──────────────────────────────────┬────────┬─────────┬─────────┬──────────┬──────────┬──────────┬──────────╮
│             Request             │              Method              │ N° Req │ Success │ Failure │   Avg    │   Max    │   Min    │   Med    │
├─────────────────────────────────┼──────────────────────────────────┼────────┼─────────┼─────────┼──────────┼──────────┼──────────┼──────────┤
│ http://localhost:8000/load/post │ POST                             │ 10     │ 10      │ 0       │ 0.002900 │ 0.004000 │ 0.001000 │ 0.003000 │
├─────────────────────────────────┼──────────────────────────────────┼────────┼─────────┼─────────┼──────────┼──────────┼──────────┼──────────┤
│ http://localhost:8000/load/put  │ PUT                              │ 10     │ 10      │ 0       │ 0.003900 │ 0.005000 │ 0.002000 │ 0.004500 │
├─────────────────────────────────┼──────────────────────────────────┼────────┼─────────┼─────────┼──────────┼──────────┼──────────┼──────────┤
│ http://localhost:8000/load/get  │ GET                              │ 10     │ 10      │ 0       │ 0.002900 │ 0.006000 │ 0.001000 │ 0.002500 │
├─────────────────────────────────┼──────────────────────────────────┼────────┼─────────┼─────────┼──────────┼──────────┼──────────┼──────────┤
│             Summary             │ Total elapsed time: 30.04329124s │        │         │         │          │          │          │          │
╰─────────────────────────────────┴──────────────────────────────────┴────────┴─────────┴─────────┴──────────┴──────────┴──────────┴──────────╯
```


## 🛠️ Configuration
###  spec.api example

Here is an example with 3 routes, but you can have as many as you want
```
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
```

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


## 👤 Authors

- **Raphael Obadia** 