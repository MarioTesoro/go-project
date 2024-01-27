# TO RUN
export GO111MODULE=on
templ generate view
go run .

# TO BUILD
go build

# TEMPL + GIN
https://github.com/a-h/templ/blob/main/examples/integration-gin/

# HTMX

# /api/register
curl -POST -d '{"username":"user","password":"password"}' http://localhost:8080/api/register
# /api/login
curl -POST -d '{"username":"user","password":"password"}' http://localhost:8080/api/login
# /api/admin/user
curl -H "Authorization: token eyJh..." http://localhost:8080/api/admin/user

