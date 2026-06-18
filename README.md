# Bolão da Copa

Aplicação de bolão da Copa do Mundo desenvolvida em Go para estudo de arquitetura de software, integração com APIs externas, persistência de dados e renderização server-side.

## Tecnologias

- Go
- Gin
- PostgreSQL
- pgx
- golang-migrate
- Templ
- HTMX
- Docker

## Funcionalidades

### Implementadas

- Importação de partidas da Copa do Mundo via Football Data API
- Persistência das partidas em PostgreSQL
- Migrações de banco de dados com golang-migrate
- Renderização server-side com Templ
- Interface dinâmica com HTMX

### Próximos Passos

- Cadastro de usuários
- Autenticação
- Sistema de palpites
- Ranking
- Sincronização automática das partidas
- Deploy

## Executando Localmente

### Variáveis de ambiente

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5432/worldcuppool
FOOTBALL_DATA_API_KEY=sua_chave
```

### Banco de dados

```bash
docker compose up -d
```

### Aplicação

```bash
air
```

ou

```bash
go run cmd/server/main.go
```

## Estrutura

```text
cmd/
internal/
web/

internal/
├── auth/
├── home/
├── match/
├── user/
└── database/
```

## Objetivo

Construir uma aplicação completa utilizando Go, explorando integração com APIs externas, persistência de dados, renderização server-side e deploy em produção.
