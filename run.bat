go build -o bookings.exe ./cmd/web/.
bookings.exe -dbname=bookings -dbuser=postgres -dbpass=admin1010 -cache=false -production=false