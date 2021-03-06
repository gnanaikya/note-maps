module github.com/google/note-maps

go 1.14

require (
	github.com/99designs/keyring v1.1.5
	github.com/AndreasBriese/bbloom v0.0.0-20190825152654-46b345b51c96 // indirect
	github.com/alecthomas/participle v0.5.0
	github.com/dgraph-io/badger v1.6.1
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/genjidb/genji v0.6.0
	github.com/golang/protobuf v1.4.2
	github.com/google/subcommands v1.2.0
	github.com/textileio/go-threads v0.1.23
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
)

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
