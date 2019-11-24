module github.com/growerlab/codev-svc

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191122220453-ac88ee75c92c
	golang.org/x/net => github.com/golang/net v0.0.0-20191119073136-fc4aabc6c914
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191120155948-bd437916bb0e
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190425163242-31fd60d6bfdc
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/graphql-go/graphql v0.7.8
	github.com/graphql-go/handler v0.2.3
	github.com/joho/godotenv v1.3.0
	github.com/rs/zerolog v1.15.0
	gopkg.in/libgit2/git2go.v27 v27.0.0-20190813173601-94786f9c8fbc
)
