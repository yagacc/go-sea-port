# Sea Port 

This monorepo includes all code for sea port management.

## Stack

- Client API - gRPC & REST services for external clients to manage ports
- Repository - gRPC service for persisting ports to Postgres

# Usage

A [Makefile](./Makefile) is provided to help development.

To run full suite run: `make`

You can then curl the GET endpoint: `curl -s http://localhost:18000/v1/port/ZABFN | jq`

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