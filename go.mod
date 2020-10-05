module pikaso

go 1.15

replace github.com/msoerjanto/pikaso/data => ./data

replace github.com/msoerjanto/pikaso/application => ./application

require (
	github.com/aws/aws-sdk-go v1.35.3
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-delve/delve v1.5.0 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/msoerjanto/pikaso/application v0.0.0-00010101000000-000000000000
	github.com/msoerjanto/pikaso/data v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20201002202402-0a1ea396d57c // indirect
	golang.org/x/text v0.3.3 // indirect
)
