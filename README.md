# Go Ran Some Where: Containerized C2 Project

This project demonstrates how to containerize a Python C2 server and a Go implant using Docker. The setup allows both components to run together with all dependencies managed within a Docker container.

## Project Structure

```
containerized-c2-project
├── c2_server.py        # Python C2 server implementation
├── ransomware.go       # Go implant implementation
├── Dockerfile          # Dockerfile for building the container
├── requirements.txt    # Python dependencies for the C2 server
├── go.mod              # Go module definition and dependencies
└── go.sum              # Checksums for Go dependencies
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd containerized-c2-project
   ```

2. **Build the Docker image:**
   ```
   docker build -t c2-project .
   ```

3. **Run the Docker container:**
   ```
   docker run -it --rm c2-project
   ```

## Usage Guidelines

- The `c2_server.py` file handles incoming connections and manages communication with the Go implant.
- The `ransomware.go` file defines the behavior of the ransomware and its interaction with the C2 server.
- Ensure that the necessary ports are exposed in the Dockerfile for communication between the server and the implant.

## Dependencies

- Python dependencies are listed in `requirements.txt`.
- Go dependencies are managed through `go.mod` and `go.sum`.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
