# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built in Go version 1.20
- Uses the [chi router](https://github.com/go-chi/chi/v5)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs/v2) session management
- Uses [nosurf](https://github.com/justinas/nosurf)

To run test with coverage HTML output, run the following command

go test -coverprofile=coverage.out
go tool cover -html=coverage.out