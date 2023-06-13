# Redis Queue Worker using `github.com/adjust/rmq/v5`

This repository implements a Redis queue worker using the `github.com/adjust/rmq/v5` library. The Redis message queue (RMQ) is a popular mechanism for building scalable and distributed systems, where different components can communicate asynchronously through a shared queue.

## Prerequisites

Before getting started, make sure you have the following:

- Go programming language installed (version 1.16 or higher).
- A running Redis server.

## Getting Started

To get started with this repository, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/nikhilryan/go-worker.git
```

2. Change into the project directory:

```bash
cd go-worker
```

3. Install the dependencies:

```bash
go mod download
```

4. Configure the Redis connection:

In the `main.go` file, update the Redis connection parameters as needed. The default configuration assumes Redis running locally on the default port (6379).

5. Build the executable:

```bash
go build
```

6. Run the worker:

```bash
./go-worker
```

The worker will start listening for messages in the Redis queue and process them as they arrive.

## Usage

To enqueue messages for the worker to process, you can use the Redis client library of your choice to push messages onto the configured queue. The worker will automatically fetch and process these messages.

Make sure the messages you enqueue are compatible with the worker's message processing logic.

## Configuration

The `main.go` file contains the configuration for the worker. You can adjust the following parameters as needed:

- `RedisAddress`: The address of the Redis server.
- `RedisPassword`: The password for connecting to the Redis server (if required).
- `RedisDB`: The Redis database to use.
- `QueueName`: The name of the Redis queue to listen to.
- `Concurrency`: The number of worker goroutines to run concurrently.

## Contributing

Contributions to this repository are welcome! If you encounter any issues or have ideas for improvements, please open an issue or submit a pull request.

When contributing, please follow the existing code style and ensure that any changes are well-documented.

## Acknowledgments

This repository is based on the `github.com/adjust/rmq/v5` library, which provides a powerful Redis message queue implementation for Go. Special thanks to the contributors of that library for their excellent work.
