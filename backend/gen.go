package main

//go:generate tern migrate --migrations ./infra/database/pgstore/migrations --config ./infra/database/pgstore/migrations
//go:generate sqlc generate -f ./infra/database/pgstore/sqlc.yaml
