## Request sequence

```mermaid
sequenceDiagram
    autonumber
    participant client as Client
    participant kong as Kong Proxy
    participant plugin_access as Plugin:access
    participant plugin_response as Plugin:response
    participant irs as Idempotent Request
    participant mongo as MongoDB
    participant api as API Server
    
    client ->> +kong: HTTP POST /*, "Idempotency-Key: $(value)"
    kong ->> +plugin_access: Trigger :access()
    plugin_access ->> +irs: HTTP PUT /captures
    irs ->> +mongo: Insert, if not exists

    alt Another request was completed
        mongo -->> irs: status: completed
        irs -->> plugin_access: 200 Ok
        plugin_access -->> kong: Set kong.response and exit
        kong -->> client: Response
    end
    
    alt Another request is in-flight
        mongo -->> irs: status: capture_is_in_flight
        irs -->> plugin_access: 409 Conflict
        plugin_access -->> kong: Set kong.response and exit
        kong -->> client: 409 Conflict
    end
    
    mongo -->> -irs: allocated
    irs -->> -plugin_access: 202 Accepted
    plugin_access -->> -kong: :access() finished
    
    kong ->> +api: Proxy the request
    api -->> -kong: Response
    
    kong ->> +plugin_response: Trigger :response()
    plugin_response ->> +irs: HTTP POST /captures
    irs ->> +mongo: Update the capture document, iff status is "allocated"
    mongo -->> -irs: Ok
    irs -->> -plugin_response: 200 Ok
    plugin_response -->> -kong: :response() completed
    
    kong -->> -client: Response
    
```