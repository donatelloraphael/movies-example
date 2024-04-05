# Setup

- Clone this repository to your local filesystem and open a terminal in that directory.
- Install Go and and setup a local PostgreSQL installation and create a DB with User and Password on Postgres.
- Give the database credentials in the line 45 of `cmd/api/main.go` file replacing the credentials in the format given.
- Install `goose` globally to handle DB migrations by using this command `go install github.com/pressly/goose/v3/cmd/goose@latest`
- install `air` globally to handle hot reload by running the command `go install github.com/cosmtrek/air@latest`
- To run the migrations, first make sure Postgres is running and run the following command in the repo root directory: `goose -dir ./migrations postgres 'user={YOUR_USERNAME_FOR_POSTGRES} password={PASSWORD_FOR_POSTGRES} host=localhost dbname={DATABASE_NAME} sslmode=disable' up`
- To run the application, run the command `air` in the repo root directory.


# Tips

- The API routes are defined in the file `cmd/api/routes.go` file.
- To create any new migrations, run the command `goose -dir ./migrations postgres 'user={YOUR_USERNAME_FOR_POSTGRES} password={PASSWORD_FOR_POSTGRES} host=localhost dbname={DATABASE_NAME} sslmode=disable' create {your-migration-name} sql`

# Exercise

- Create an API endpoint for updating the name of the user with the provided email. Remember, only the logged in user must be allowed to change his/her name. 
- You may need to create a middleware to check whether the user is logged in or not. You can refer to the completed project at [https://github.com/donatelloraphael/greenlight](https://github.com/donatelloraphael/greenlight) if you get stuck. It also provides examples of couple of functionalities like middelwares which we did not talk about in the training session.
