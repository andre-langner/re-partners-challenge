# RE partners

Software Engineering Challenge

## Dependencies

* [Go 1.21+](https://go.dev)
* [Docker](https://www.docker.com)
* [Make](https://www.gnu.org/software/make/)

## Architecture

The architecture is based on the *Clean Architecture* design, also it was build following the *S.O.L.I.D.* principles.

This choice makes it possible to separate layer responsibilities and isolate external dependencies, enabling easy testing and facilitating the implementation of new functionality.

## Testing

Into the root folder run the following command.
There's no need to have Go installed [Containerized testing].
```sh
make test
```

## API Documentation

The documentation regarding the endpoint can be found in the *docs/openapi.yaml* file.

## Running

Into the root folder run the following command. The API uses by default the port 3000.

```sh
make run
```

