

## Endpoints da API

### Usuários (`/api/users`)

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/api/users` | Criar novo usuário |
| GET | `/api/users/by-email?email=` | Buscar usuário por email |
| GET | `/api/users` | Listar todos os usuários |
| GET | `/api/users/:id` | Buscar usuário por ID |
| PUT | `/api/users/:id` | Atualizar usuário |
| DELETE | `/api/users/:id` | Deletar usuário |

### Livros (`/api/books`)

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/api/books` | Criar novo livro |
| GET | `/api/books/by-isbn?isbn=` | Buscar livro por ISBN |
| GET | `/api/books` | Listar todos os livros |
| GET | `/api/books/:id` | Buscar livro por ID |
| PUT | `/api/books/:id` | Atualizar livro |
| DELETE | `/api/books/:id` | Deletar livro |

### Empréstimos (`/api/loans`)

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/api/loans` | Criar novo empréstimo |
| GET | `/api/loans/:id` | Buscar empréstimo por ID |
| PUT | `/api/loans/:id/return` | Devolver livro |
| GET | `/api/loans` | Listar todos os empréstimos |
| GET | `/api/loans/by-user/:user_id` | Listar empréstimos por usuário |
| GET | `/api/loans/by-book/:book_id` | Listar empréstimos por livro |
| DELETE | `/api/loans/:id` | Deletar empréstimo |