# Documentation

Stark Bank backend coding challenge

## Test Requirements
- Issue 8 to 12 invoices to random clients every 3h for 24h
- Receive webhook callbacks of invoice changes
- Tranfer paid value, minus fees, to provided account.  

## Project minimum
- Go 1.25+
- Deploy requirements:
    - AWS account and CLI configured
    - Terrafrom CLI
    - Repository secrets for GH Action

## Instalation and running
```bash
go mod tidy
```
### Configure .env

```env
PROJECT_ID="project_id"

PRIVATE_KEY="generated_key"

# Porvided values
BANK_CODE=""
BRANCH=""
ACCOUNT=""
NAME=""
TAX_ID=""
ACCOUNT_TYPE=""
```

```bash
go run main.go
```

## Core functionalitty

```golang
server.POST("/start-cron", app.StartCron)
```
Starts a cron job that runs every 3h for 24h, that generates 8 to 12 invoices and calls the stark bank api via SDK and issues the generated invoices.

It runs only one job at a time.

```golang
server.POST("/webhook/payment", app.Listener)
```
Receives the webhook callback. Uses SDK parses to validate signature, and parse the josn body to a struct, and transfers only if invoice is credited.

### Custom parser reason

The SDK is currently returning a generic interface, so I decided to get the body from request and parse to a struct based on json example from the documentation.

## Errors

Errors are handled in a way that it doesn't kill the app, but are registred in a txt log.

## Dashboard (easier to see deployed app running)

Loads the logging to a json file to show on frontend.
```golang
server.GET("/", app.DashboardHandler) //Visual
server.GET("/data", app.DashboardDataHandler) // Raw json
```

## Tests

Implemented a few unit tests:

- Validate if generator returns 10 invoices
- Validate if client name and cpf matches
- Validate a if parses valid json
- Validate if invalid json is handed correctly
- Validate if an invalid signature is treated correctly
- Validate if empty body is handled before calling the parser

## Deploy

Application is deployed to AWS via GH Actions. The action builds and upload binary to EC2 instance, and run application on background.

Terraform is used to handle building, changes and destroying AWS infra safely. 