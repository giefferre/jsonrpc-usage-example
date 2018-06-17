# JSON-RPC Usage Example

A simple client-server architecture showing how JSON-RPC 2.0 works.

Contains:

- A Server, implemented in Go
- A Client, implemented in Python

## Running

In order to simplify the execution, a convenient Makefile is provided.

That said, [Docker](https://www.docker.com/community-edition) and `make` is all you need to run both server and client.

### Server

Run server first by opening a bash console and executing `make run-server`.

The command will download Go image, get the dependencies, build the application and run it.

### Client

Open a bash console and execute `make run-client`.

The command will download Python 3 image, get dependencies and run the application.