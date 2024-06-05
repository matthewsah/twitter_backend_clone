# twitter_backend_clone

## TO RUN

### Create the database in PostgreSQL
Use CLI, pgAdmin, or another tool to create the twitter_clone_db

### Perform database migrations if necessary
`cd sql/schema`
`psql -h localhost -U <username> -l`
`createdb -h localhost -U <username> twitter_clone_db`
`goose postgres <DB_URL> up`

### Build the code
`go build`

### Run
`./twitter_backend_clone`

### To make calls
localhost:8080/v1/{YOUR_API_CALL}