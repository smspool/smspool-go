package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	smspool "github.com/smspool/smspool-go"
)

func fail(msg string, args ...any) { fmt.Fprintf(os.Stderr, msg+"\n", args...); os.Exit(1) }

func main() {
	cfg := smspool.NewConfiguration()
	client := smspool.NewAPIClient(cfg)
	ctx := context.Background()

	countries, _, err := client.InformationAPI.CountryRetrieveAll(ctx).Execute()
	if err != nil { fail("countries: %v", err) }
	list, ok := countries.([]interface{})
	if !ok || len(list) == 0 { fail("no countries: %#v", countries) }
	first, _ := json.Marshal(list[0])
	fmt.Println("countries:", len(list), string(first))

	rates, _, err := client.InformationAPI.RequestSuccessRate(ctx).Service("1").Execute()
	if err != nil || rates == nil { fail("success rate: %v", err) }
	b, _ := json.Marshal(rates)
	fmt.Println("success_rate (form POST):", string(b)[:80])

	authed := context.WithValue(ctx, smspool.ContextAccessToken, strings.Repeat("0", 32))
	_, resp, err := client.InformationAPI.RequestBalance(authed).Execute()
	if err == nil || resp == nil || (resp.StatusCode != 401 && resp.StatusCode != 403) { fail("expected an auth error, got %v %v", err, resp) }
	var apiErr *smspool.GenericOpenAPIError
	if e, ok := err.(*smspool.GenericOpenAPIError); ok { apiErr = e }
	body := ""
	if apiErr != nil { body = string(apiErr.Body()) }
	if len(body) > 80 { body = body[:80] }
	fmt.Println("bad key ->", resp.StatusCode, body)
	fmt.Println("GO OK")
}
