# Mini E-Wallet Service

## Description
This service covers some of the essential REST APIs of an E-Wallet application including authentication system and transaction management system.

## Installation
1. Run `make setup` command to make a copy of config example, fetch all dependencies, and install swagger command.
2. Update `internal/app/config.yaml` file with your own credentials. <br/> NOTE: If you change You need to change `server.address` and/or `server.port.http`, please update swagger annotations in `main.go` as well.
3. Create new database under the same schema `as database.sql.schema` in config file.
4. Run `make run` to generate API docs and start the http server. Or simply run `go run main.go` to start the server without re-generating API docs.
5. Open `http://localhost:2201/swagger/index.html` in your browser to see the contracts.

## License
[MIT](https://github.com/ffauzann/mini-e-wallet/blob/main/LICENSE)
