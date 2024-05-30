# go-pod

A simple application to write and run smoke tests for RESTful APIs.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

### Steps

1. Clone the repository:

    ```bash
    git clone https://github.com/jdnielss/go-pod.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-pod
    ```

3. Navigate to the project directory:

    ```bash
    go mod tidy
    ```

4. Build the project:

    ```bash
    make build
    ```

### Running

The most convenient way of running this code, especially in a CI environment, is to use the docker image `jdnielss/pod`

`docker run --rm -v "$(pwd)":/test jdnielss/pod -f /test/config.yaml -h http://{YOUR_URL}`


## Usage

Explain how to use your project. Provide examples if applicable.

```bash
./pod -f config.yaml -h http://localhost:8080
```

- `-f` flag: Path to the YAML config file.
- `-h` flag: Base URL of the API.

## Configuration

Describe the configuration options and how to customize them.

## Contributing

If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/yourfeature`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add some feature'`).
5. Push to the branch (`git push origin feature/yourfeature`).
6. Create a new Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).
