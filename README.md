# Idempotent Requests
Idempotent Requests Server provides API to allocate and record a captured request.

## Project status

 ![CI workflow](https://github.com/checkr/idempotent-requests/actions/workflows/ci.yml/badge.svg)
 ![Release workflow](https://github.com/checkr/idempotent-requests/actions/workflows/release.yml/badge.svg)

## Configuration

| Config | Required | Default | Explanation |
| --- | --- | --- | --- |
| `MONGODB_URI` | yes | `mongodb://root:password123@localhost:27017` | URI to connect to MongoDB |


## Example request sequence

![sequence](./docs/sequence.png)

## Running Integration Tests Locally

```shell
docker-compose -f docker-compose.yml \
  -f docker-compose.server.yml \
  -f docker-compose.integration.yml \
  up --exit-code-from integration-test
```