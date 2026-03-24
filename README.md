# auth-service-go-postgres

Production-ready backend starter (Senior-level baseline).

## Stack
- Language: go
- Database: postgres

## Included
- Layered structure starter
- Docker + docker-compose (app + db)
- .env.example
- Basic tests scaffold
- GitHub Actions CI

## Run
`ash
docker compose up -d
`

## Next hardening
- Add real DB migrations
- Add integration tests against DB container
- Add auth/rate-limit and tracing
