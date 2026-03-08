# Documentation

Stark Bank backend coding challenge

## Core functionalitty

```golang
server.POST("/start-cron", app.StartCron)
```
Starts a cron job that runs every 3h for 24h, that generates 10 invoices and calls the stark bank api via SDK and issues the generated invoices

```golang
server.POST("/webhook/payment", app.Listener)
```
Receives the webhook callback. Uses SDK parses to validate signature, and parse the josn body to a struct, and transfers only if invoice is credited.

### Custom parses reason

The SDK is currently returning a generic interface, so I decided to get the body from request and parse to a struct based on json example from the documentation.

## Tests

Implemented a few unit tests:

- Validate if generator returns 10 invoices
- Validate if client name and cpf matches
- Validate a if parses valid json
- Validate if invalid json is handed correctly
- Validate if an invalid signature is treated correctly
- Validate if empty body is handled before calling the parser

## Deploy

Application is deployed to AWS via GH Actions. The action builds and upload binary to EC2 instance.

Terraform is used to handle builds, changes and destroying infra safely. 