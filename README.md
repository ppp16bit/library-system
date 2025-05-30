# Como Rodar o Container

**Pré-requisito:** Docker instalado, PostgreSQL 16

## 1. Configurar Variáveis

Crie um arquivo `.env` na raiz do projeto:

```env
DB_NAME=
DB_USER=
DB_PASSWORD=
DB_PORT=
PORT_APP=
```

## 2. Executar

```bash
# Primeira vez (com build)
docker-compose up --build

# Próximas vezes
docker-compose up

# Em background
docker-compose up -d
```

## 3. Parar

```bash
docker-compose down
```

## Acessos

- **Backend:** http://localhost:8080
- **PostgreSQL:** localhost:5433