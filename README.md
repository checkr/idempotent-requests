![CI workflow](https://github.com/checkr/idempotent-requests/actions/workflows/ci.yml/badge.svg)
![Release workflow](https://github.com/checkr/idempotent-requests/actions/workflows/release.yml/badge.svg)

# Idempotent Requests

Idempotent Requests Server provides [API](./docs/openapi.yaml) to allocate and record a captured request. It follows
behaviors described
in [IETF draft for standardized `Idempotency-Key` header](https://datatracker.ietf.org/doc/html/draft-ietf-httpapi-idempotency-key-header-00)
.

## Contract

[API contract](./docs/openapi.yaml)

## MongoDB access pattern

[More details on MongoDB access pattern](./docs/mongo_upsert.md)

## Idempotent Requests Client

We have implemented a client in a form of a [Kong Gateway](https://konghq.com/kong/)
Plugin - [`idempotent-requests`](https://github.com/checkr/kong-plugin-idempotent-requests).

## Configuration

| Config | Required | Default | Explanation |
| --- | --- | --- | --- |
| `MONGODB_URI` | yes | `mongodb://root:password123@localhost:27017` | URI to connect to MongoDB |

## Example topology

This example topology relies on Kong Gateway to intercept client requests.
A [Kong Plugin](https://github.com/checkr/kong-plugin-idempotent-requests) acts as a client to Idempotent Requests
Server.

### Data flow

![data_flow](./docs/example_data_flow.png)

### Sequence

[Source](./docs/sequence.md)

[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgYXV0b251bWJlclxuICAgIHBhcnRpY2lwYW50IGNsaWVudCBhcyBDbGllbnRcbiAgICBwYXJ0aWNpcGFudCBrb25nIGFzIEtvbmcgUHJveHlcbiAgICBwYXJ0aWNpcGFudCBwbHVnaW5fYWNjZXNzIGFzIFBsdWdpbjphY2Nlc3NcbiAgICBwYXJ0aWNpcGFudCBwbHVnaW5fcmVzcG9uc2UgYXMgUGx1Z2luOnJlc3BvbnNlXG4gICAgcGFydGljaXBhbnQgaXJzIGFzIElkZW1wb3RlbnQgUmVxdWVzdFxuICAgIHBhcnRpY2lwYW50IG1vbmdvIGFzIE1vbmdvREJcbiAgICBwYXJ0aWNpcGFudCBhcGkgYXMgQVBJIFNlcnZlclxuICAgIFxuICAgIGNsaWVudCAtPj4gK2tvbmc6IEhUVFAgUE9TVCAvKiwgXCJJZGVtcG90ZW5jeS1LZXk6ICQodmFsdWUpXCJcbiAgICBrb25nIC0-PiArcGx1Z2luX2FjY2VzczogVHJpZ2dlciA6YWNjZXNzKClcbiAgICBwbHVnaW5fYWNjZXNzIC0-PiAraXJzOiBIVFRQIFBVVCAvY2FwdHVyZXNcbiAgICBpcnMgLT4-ICttb25nbzogSW5zZXJ0LCBpZiBub3QgZXhpc3RzXG5cbiAgICBhbHQgQW5vdGhlciByZXF1ZXN0IHdhcyBjb21wbGV0ZWRcbiAgICAgICAgbW9uZ28gLS0-PiBpcnM6IHN0YXR1czogY29tcGxldGVkXG4gICAgICAgIGlycyAtLT4-IHBsdWdpbl9hY2Nlc3M6IDIwMCBPa1xuICAgICAgICBwbHVnaW5fYWNjZXNzIC0tPj4ga29uZzogU2V0IGtvbmcucmVzcG9uc2UgYW5kIGV4aXRcbiAgICAgICAga29uZyAtLT4-IGNsaWVudDogUmVzcG9uc2VcbiAgICBlbmRcbiAgICBcbiAgICBhbHQgQW5vdGhlciByZXF1ZXN0IGlzIGluLWZsaWdodFxuICAgICAgICBtb25nbyAtLT4-IGlyczogc3RhdHVzOiBjYXB0dXJlX2lzX2luX2ZsaWdodFxuICAgICAgICBpcnMgLS0-PiBwbHVnaW5fYWNjZXNzOiA0MDkgQ29uZmxpY3RcbiAgICAgICAgcGx1Z2luX2FjY2VzcyAtLT4-IGtvbmc6IFNldCBrb25nLnJlc3BvbnNlIGFuZCBleGl0XG4gICAgICAgIGtvbmcgLS0-PiBjbGllbnQ6IDQwOSBDb25mbGljdFxuICAgIGVuZFxuICAgIFxuICAgIG1vbmdvIC0tPj4gLWlyczogYWxsb2NhdGVkXG4gICAgaXJzIC0tPj4gLXBsdWdpbl9hY2Nlc3M6IDIwMiBBY2NlcHRlZFxuICAgIHBsdWdpbl9hY2Nlc3MgLS0-PiAta29uZzogOmFjY2VzcygpIGZpbmlzaGVkXG4gICAgXG4gICAga29uZyAtPj4gK2FwaTogUHJveHkgdGhlIHJlcXVlc3RcbiAgICBhcGkgLS0-PiAta29uZzogUmVzcG9uc2VcbiAgICBcbiAgICBrb25nIC0-PiArcGx1Z2luX3Jlc3BvbnNlOiBUcmlnZ2VyIDpyZXNwb25zZSgpXG4gICAgcGx1Z2luX3Jlc3BvbnNlIC0-PiAraXJzOiBIVFRQIFBPU1QgL2NhcHR1cmVzXG4gICAgaXJzIC0-PiArbW9uZ286IFVwZGF0ZSB0aGUgY2FwdHVyZSBkb2N1bWVudCwgaWZmIHN0YXR1cyBpcyBcImFsbG9jYXRlZFwiXG4gICAgbW9uZ28gLS0-PiAtaXJzOiBPa1xuICAgIGlycyAtLT4-IC1wbHVnaW5fcmVzcG9uc2U6IDIwMCBPa1xuICAgIHBsdWdpbl9yZXNwb25zZSAtLT4-IC1rb25nOiA6cmVzcG9uc2UoKSBjb21wbGV0ZWRcbiAgICBcbiAgICBrb25nIC0tPj4gLWNsaWVudDogUmVzcG9uc2UiLCJtZXJtYWlkIjp7InRoZW1lIjoiZGVmYXVsdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlLCJhdXRvU3luYyI6dHJ1ZSwidXBkYXRlRGlhZ3JhbSI6ZmFsc2V9)](https://mermaid.live/edit#eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgYXV0b251bWJlclxuICAgIHBhcnRpY2lwYW50IGNsaWVudCBhcyBDbGllbnRcbiAgICBwYXJ0aWNpcGFudCBrb25nIGFzIEtvbmcgUHJveHlcbiAgICBwYXJ0aWNpcGFudCBwbHVnaW5fYWNjZXNzIGFzIFBsdWdpbjphY2Nlc3NcbiAgICBwYXJ0aWNpcGFudCBwbHVnaW5fcmVzcG9uc2UgYXMgUGx1Z2luOnJlc3BvbnNlXG4gICAgcGFydGljaXBhbnQgaXJzIGFzIElkZW1wb3RlbnQgUmVxdWVzdFxuICAgIHBhcnRpY2lwYW50IG1vbmdvIGFzIE1vbmdvREJcbiAgICBwYXJ0aWNpcGFudCBhcGkgYXMgQVBJIFNlcnZlclxuICAgIFxuICAgIGNsaWVudCAtPj4gK2tvbmc6IEhUVFAgUE9TVCAvKiwgXCJJZGVtcG90ZW5jeS1LZXk6ICQodmFsdWUpXCJcbiAgICBrb25nIC0-PiArcGx1Z2luX2FjY2VzczogVHJpZ2dlciA6YWNjZXNzKClcbiAgICBwbHVnaW5fYWNjZXNzIC0-PiAraXJzOiBIVFRQIFBVVCAvY2FwdHVyZXNcbiAgICBpcnMgLT4-ICttb25nbzogSW5zZXJ0LCBpZiBub3QgZXhpc3RzXG5cbiAgICBhbHQgQW5vdGhlciByZXF1ZXN0IHdhcyBjb21wbGV0ZWRcbiAgICAgICAgbW9uZ28gLS0-PiBpcnM6IHN0YXR1czogY29tcGxldGVkXG4gICAgICAgIGlycyAtLT4-IHBsdWdpbl9hY2Nlc3M6IDIwMCBPa1xuICAgICAgICBwbHVnaW5fYWNjZXNzIC0tPj4ga29uZzogU2V0IGtvbmcucmVzcG9uc2UgYW5kIGV4aXRcbiAgICAgICAga29uZyAtLT4-IGNsaWVudDogUmVzcG9uc2VcbiAgICBlbmRcbiAgICBcbiAgICBhbHQgQW5vdGhlciByZXF1ZXN0IGlzIGluLWZsaWdodFxuICAgICAgICBtb25nbyAtLT4-IGlyczogc3RhdHVzOiBjYXB0dXJlX2lzX2luX2ZsaWdodFxuICAgICAgICBpcnMgLS0-PiBwbHVnaW5fYWNjZXNzOiA0MDkgQ29uZmxpY3RcbiAgICAgICAgcGx1Z2luX2FjY2VzcyAtLT4-IGtvbmc6IFNldCBrb25nLnJlc3BvbnNlIGFuZCBleGl0XG4gICAgICAgIGtvbmcgLS0-PiBjbGllbnQ6IDQwOSBDb25mbGljdFxuICAgIGVuZFxuICAgIFxuICAgIG1vbmdvIC0tPj4gLWlyczogYWxsb2NhdGVkXG4gICAgaXJzIC0tPj4gLXBsdWdpbl9hY2Nlc3M6IDIwMiBBY2NlcHRlZFxuICAgIHBsdWdpbl9hY2Nlc3MgLS0-PiAta29uZzogOmFjY2VzcygpIGZpbmlzaGVkXG4gICAgXG4gICAga29uZyAtPj4gK2FwaTogUHJveHkgdGhlIHJlcXVlc3RcbiAgICBhcGkgLS0-PiAta29uZzogUmVzcG9uc2VcbiAgICBcbiAgICBrb25nIC0-PiArcGx1Z2luX3Jlc3BvbnNlOiBUcmlnZ2VyIDpyZXNwb25zZSgpXG4gICAgcGx1Z2luX3Jlc3BvbnNlIC0-PiAraXJzOiBIVFRQIFBPU1QgL2NhcHR1cmVzXG4gICAgaXJzIC0-PiArbW9uZ286IFVwZGF0ZSB0aGUgY2FwdHVyZSBkb2N1bWVudCwgaWZmIHN0YXR1cyBpcyBcImFsbG9jYXRlZFwiXG4gICAgbW9uZ28gLS0-PiAtaXJzOiBPa1xuICAgIGlycyAtLT4-IC1wbHVnaW5fcmVzcG9uc2U6IDIwMCBPa1xuICAgIHBsdWdpbl9yZXNwb25zZSAtLT4-IC1rb25nOiA6cmVzcG9uc2UoKSBjb21wbGV0ZWRcbiAgICBcbiAgICBrb25nIC0tPj4gLWNsaWVudDogUmVzcG9uc2UiLCJtZXJtYWlkIjoie1xuICBcInRoZW1lXCI6IFwiZGVmYXVsdFwiXG59IiwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ)

## Running Integration Tests Locally

```shell
docker-compose -f docker-compose.yml \
  -f docker-compose.server.yml \
  -f docker-compose.integration.yml \
  up --exit-code-from integration-test
```