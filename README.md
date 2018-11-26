DB:

for using heroku db export env variable
DB_URI=postgres://qjatviiirqkyco:bb2c48f1a38452ce6806e8c71e4001fa74937be441845667ac2046aa7e267fbd@ec2-54-75-231-3.eu-west-1.compute.amazonaws.com:5432/d9hf716191g6j0

================

Start Server:

go run main.go

================

API (more information about endpoints in ./routers/routers.go)

PUBLIC

POST /registrations  - save user in system

POST /authorizations - get token for protected api

PROTECTED

GET /films - get film's list

