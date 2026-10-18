# SMSPool Go SDK

Official Go client for the [SMSPool API](https://api.smspool.net/docs): one-time SMS
verification numbers, rentals, pre-orders, eSIMs and more.

Package: [`github.com/smspool/smspool-go`](https://pkg.go.dev/github.com/smspool/smspool-go) on the Go module proxy.

## Install

```bash
go get github.com/smspool/smspool-go
```

## Quick start

Get your API key from [your settings page](https://www.smspool.net/my/settings).

```go
import smspool "github.com/smspool/smspool-go"

client := smspool.NewAPIClient(smspool.NewConfiguration())
ctx := context.WithValue(context.Background(), smspool.ContextAccessToken, "YOUR_API_KEY")
balance, _, err := client.InformationAPI.RequestBalance(ctx).Execute()
countries, _, err := client.InformationAPI.CountryRetrieveAll(ctx).Execute()
order, _, err := client.SMSAPI.PurchaseSms(ctx).Country("US").Service("1").Execute()
```

Responses are decoded JSON (`map[string]interface{}` or `[]interface{}`). Errors (HTTP 4xx/5xx)
are `*smspool.GenericOpenAPIError`; `Body()` holds the response.

## Calling endpoints

Every endpoint in the [API reference](https://api.smspool.net/docs) is a method named after
its path, on the class for its section: `/sms/check` is `SmsCheck()` on `client.SMSAPI`. Endpoints that need
your API key send it as a bearer token.

Responses are returned as decoded JSON exactly as the API sends them: many endpoints answer
`200` with `"success": 0` and a `message` when a request is refused, so check `success`.

- Generated reference for every method and parameter: [`docs/`](docs/)
- API reference with examples: https://api.smspool.net/docs

## How this SDK is made

The code is generated from the API's OpenAPI spec
([openapi.json](https://api.smspool.net/openapi.json)) with
[OpenAPI Generator](https://openapi-generator.tech). Do not edit generated files by hand:
a workflow regenerates them every week and opens a pull request when the API changed.
Hand-written files are listed in `.openapi-generator-ignore`.

## Releasing (maintainers)

1. Set the new version in `openapi-generator.yaml` (`packageVersion`).
2. Run `scripts/generate.sh` (needs Docker) and `scripts/smoke.sh`, then commit.
3. Tag and push: `git tag vX.Y.Z && git push origin vX.Y.Z`.

The Release workflow checks the tag against the version, runs the smoke test and publishes.
There is no publish step: a Go version is the git tag. The workflow asks the module
proxy to fetch it so it shows on pkg.go.dev right away.

## License

MIT
