# KernelPanic Blog
___

## Run docker environment

```bash
docker-compose -p kernelpanic_blog --env-file ../.env up -d
```

## Configuration

Add __UUID__ extension for postgres

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

## Run Tests

```bash
go test -v ./tests
```