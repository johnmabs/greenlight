# Greenlight API

REST API développée en Go dans le cadre de mon apprentissage
du backend engineering avec Go et PostgreSQL.

L'objectif principal du projet est de construire une API
HTTP robuste en utilisant autant que possible les outils
standards de l'écosystème Go.

## Stack

- Go
- PostgreSQL
- SQL migrations
- net/http
- Makefile

## Structure

```text
.
├── cmd/
│   └── api/
├── internal/
├── migrations/
├── Makefile
├── go.mod
└── go.sum
```

## Concepts travaillés
- REST API design
- HTTP handlers
- Middleware
- JSON serialization
- Validation
- PostgreSQL
- Database migrations
- Configuration
- Error handling
- Graceful shutdown
- Go project structure

## Démarrage
```
go mod download
```
Configurer la base PostgreSQL puis lancer :
```
go run ./cmd/api
```

## Objectif du repository

Ce projet est principalement un laboratoire backend destiné
à approfondir Go, HTTP et les bonnes pratiques de conception
d'API.

**Description :**

> REST API learning project built with Go and PostgreSQL.

Topics :

```text
golang
go
postgresql
rest-api
backend
api
```


