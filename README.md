Here is the link to the frontend [repo](https://github.com/Arshad-Siddiqui/whereiapplied)

## Installation

```bash
go mod tidy
```

## Run the app

```bash
go run main.go
```

## Create a dotenv file

Create a .env file in the root of the project and add the following:

```
MONGO_URI=<your_mongoDB_Atlas_uri_with_credentials>
```

> For tests create .env.test and put in the same MONGO_URI but with a local database.
