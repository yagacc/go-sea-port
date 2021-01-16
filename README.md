# Sea Port 

This monorepo includes all code for sea port management.

## Stack

- Client API - gRPC & REST services for external clients to manage ports
- Repository - gRPC service for persisting ports to Postgres

# Usage

A [Makefile](./Makefile) is provided to help development.

To run full suite run: `make`

You will get a REST endpoint at: http://localhost:18000/health