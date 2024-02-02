# twitter_backend_clone

## TO RUN

### Create the database in PostgreSQL
Use CLI, pgAdmin, or another tool to create the twitter_clone_db

### Perform database migrations
`cd sql/schema`
`goose postgres postgres_link up`

### Build the code
`go build`

### Run
`./twitter_backend_clone`

### To make calls
localhost:8080/v1/{YOUR_API_CALL}