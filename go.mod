module resk

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/segmentio/ksuid v1.0.2
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/tietang/dbx v1.0.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e
	golang.org/x/text v0.3.0
	golang.org/x/tools v0.0.0-20190328211700-ab21143f2384
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190403202508-8e1b8d32e692
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190403144856-b630fd6fe46b
	golang.org/x/sys v0.0.0-20190403152447-81d4e9dc473e => github.com/golang/sys v0.0.0-20190403152447-81d4e9dc473e
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20190328211700-ab21143f2384 => github.com/golang/tools v0.0.0-20190404132500-923d25813098
)
