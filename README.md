# Simple Golang datastore PoC

Read through the documentation at https://github.com/gorilla/mux to see how the request handling is set up.

Built up from PoC developed with Eugene Kuprijevics

Environment variables using https://github.com/joho/godotenv

To install depedencies:
`go get github.com/joho/godotenv`
`go get github.com/gorilla/handlers`
`go get github.com/gorilla/mux`

To deploy to GCP Appengine Standard (from a locally gcloud authenticated account)/ change the GCP project in .env to your own:
`gcloud app deploy`