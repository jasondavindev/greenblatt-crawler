module github.com/jasondavindev/statusinvest-crawler

replace (
	github.com/jasondavindev/statusinvest-crawler/config => ./config
	github.com/jasondavindev/statusinvest-crawler/csv => ./csv
	github.com/jasondavindev/statusinvest-crawler/greenblatt => ./greenblatt
	github.com/jasondavindev/statusinvest-crawler/requester => ./requester
)

go 1.15

require (
	github.com/fatih/structs v1.1.0
	github.com/oleiade/reflections v1.0.1
	go.uber.org/config v1.4.0
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/tools v0.0.0-20201204222352-654352759326 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
)
