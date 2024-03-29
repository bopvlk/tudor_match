# Backend for start-up project

Backend can be available at: <https://tudor-match.fly.dev>

Swagger: <https://tudor-match.fly.dev/api/docs/index.html>

## Local run in .devcontainer 

Install extension for VS Code named Dev Conteiners

Instructions on how to use Dev Containers:
<https://code.visualstudio.com/docs/devcontainers/tutorial>


Before building project requires the presence of a .env file in .devcontainer folder (example can be found in .env.example)

```env
DB_URL= # db connection string
MIGRATIONS_URL= # db connection string for migrations in most cases it is the same as DB_URL
PORT= # port to run the server on
GOOGLE_OAUTH_CLIENT_ID= # google oauth client id
GOOGLE_OAUTH_CLIENT_SECRET= # google oauth client secret
GOOGLE_OAUTH_REDIRECT_PAGE= # google oauth redirect page domain +.../api/auth/google/callback                                   
FACEBOOK_OAUTH_CLIENT_ID= # facebook oauth client id
FACEBOOK_OAUTH_CLIENT_SECRET= # facebook oauth client secret
FACEBOOK_OAUTH_REDIRECT_PAGE= # facebook oauth redirect page domain +.../api/auth/facebook/callback
GOOGLE_EMAIL_ADDRESS= # email address to send emails from
GOOGLE_EMAIL_SECRET= # email password
PASSWORD_RESET_REDIRECT_PAGE=   # password reset redirect page domain +.../api/auth/password-reset/callback
JWT_SECRET= # jwt secret key
SQLC_AUTH_TOKEN= # sqlc auth token
COOKIE_SECRET= # cookie secret key
ALLOWED_ORIGINS= # allowed origins for cors (comma separated) for example: http://localhost:3000,http://localhost:3001
REDIRECT_URL= # redirect url for oauth (for example: http://localhost:3000)
```

Run the application locally:
```bash
make run
```

If everything is OK, you can use the application with any browser by this URL: http://localhost:_your_port_

## Documentation

Swagger requires env variable **SERVER_HOSTNAME** defined in .env file

Documentation is available at:

- <http://localhost:your_port/api/docs/index.html>
-  <https://tudor-match.fly.dev/api/docs/index.html>

 
**Dev Containers is set up to install Swagger on its own.**
But if you want to install it yourself, or you haven't installed Devcontainer run this code:
```bash
go install github.com/swaggo/swag/cmd/swag@latest`
```

- Generate docs (path to docs used in main.go: `./docs`): 
```bash
make swag
```
## DB code generation

For code generation from SQL to golang we use ***sqlc***. Documentation https://docs.sqlc.dev/en/v1.23.0/

***Dev Containers is configured to install sqlc itself*** as well.
If you want do it by yourself:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Create SQL queries in: `./database/queries/{entity_name}.sql`

In root folder with present `./sqlc.yaml` file run:
```bash
make db
```
Generated code will be in `./database/queries/`

## Migrations

Database migrations written in Go. https://github.com/golang-migrate/migrate/v4

MigrationCLI instaled in DevContainer, but if you watn do it by yourself:

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

For migration file creation use:
```bash
make migrate-create name={action_name}
```

Migration is started when you run the program with the command `make run`, howewer if you want do it by yourself:
```bash
make migration
```

Migration using files from: ``./database/migration``

## Request examples

User creation request example:

```bash
curl -v  POST http://localhost:8000/api/auth/register -H "Content-type: application/json" -d '{"email": "hello@world", "name":"world", "password":"hello"}'
```

Response body:

```json
{
  "Data": {
    "user": {
      "id": "143a5c7a-8e4b-4081-9da8-8005ee7b2e6f",
      "name": "world",
      "email": "hello@world",
      "photo": "default.jpeg",
      "verified": false,
      "password": "$2a$10$5bMojbONIM9KR3IsqvGQEO/eMOuKpRaJ8/qrSkphnaSKNkJeXbkw2",
      "role": "user",
      "created_at": "2023-09-25T07:42:26.808Z",
      "updated_at": "2023-09-25T07:42:26.804Z"
    }
  },
  "Status": "success"
}
```

User login request example:

```bash
curl -v  POST http://localhost:8000/api/auth/login -H "Content-type: application/json" -d '{"email": "hello@world", "password":"hello"}'
```

Response body:

```json
{
  "token": "some_token"
}
```

Token usage for protected endpoint:

```bash
curl -X GET "http://localhost:8000/protected/userinfo" -H "Authorization: Bearer <token generated by login endpoint>"
```

Response body:

```json
{
  "user": {
    "id": "5fd55515-f611-4182-808e-f80ca523d3c3",
    "name": "world",
    "email": "hello@world",
    "photo": "default.jpeg",
    "verified": false,
    "password": "$2a$10$CqAghRSTwM3vNtGby4aDoOIf4ezGuJA4oUzNEV4oqgic3ORN9.RM2",
    "role": "user",
    "created_at": "2023-09-25T08:28:42.302Z",
    "updated_at": "2023-09-25T08:28:42.3Z"
  }
}
```

## API diagrams

![API Diagram](./api.drawio.svg)
