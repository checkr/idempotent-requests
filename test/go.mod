module checkr.com/idempotent-requests/tests/integation

go 1.17

require checkr.com/idempotent-requests v0.1.0

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/caarlos0/env/v6 v6.7.1
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace checkr.com/idempotent-requests => ./..
