![CI workflow](https://github.com/checkr/idempotent-requests/actions/workflows/ci.yml/badge.svg)
![Release workflow](https://github.com/checkr/idempotent-requests/actions/workflows/release.yml/badge.svg)

# Idempotent Requests
Idempotent Requests Server provides API to allocate and record a captured request.

## Idempotent Requests Client

We have implemented a client in a form of a [Kong Gateway](https://konghq.com/kong/) Plugin - [`idempotent-requests`](https://github.com/checkr/kong-plugin-idempotent-requests). 

## Configuration

| Config | Required | Default | Explanation |
| --- | --- | --- | --- |
| `MONGODB_URI` | yes | `mongodb://root:password123@localhost:27017` | URI to connect to MongoDB |


## Example topology

This example topology relies on Kong Gateway to intercept client requests.
A Kong Plugin acts as a client to Idempotent Requests Server.

### Data flow 
![data_flow](./docs/example_data_flow.png)

### Sequence
![sequence](./docs/sequence.png)

## Running Integration Tests Locally

```shell
docker-compose -f docker-compose.yml \
  -f docker-compose.server.yml \
  -f docker-compose.integration.yml \
  up --exit-code-from integration-test
```