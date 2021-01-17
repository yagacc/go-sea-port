# Sea Port 

This monorepo includes all code for sea port management.

## Stack

- Client API - gRPC & REST services for external clients to manage ports
- Repository - gRPC service for persisting ports to Postgres 
- A [Makefile](./Makefile) is provided

## Init

- `git clone git@github.com:yagacc/go-sea-port.git`
- copy ports.json to private localhost directory at project root: `.localhost/spec`. E.g. ` mkdir -p .localhost/spec && cp ports.json .localhost/spec`
- run tests: `make test`
- start application: `make`
- manual test: curl GET endpoint: `curl -s http://localhost:18000/v1/port/ZABFN | jq`

  You should see something similar:
  ```json
    {
      "port": {
        "id": "ZABFN",
        "name": "Bloemfontein",
        "city": "Bloemfontein",
        "country": "South Africa"
      }
    }
  ```