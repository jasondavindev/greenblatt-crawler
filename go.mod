module github.com/jasondavindev/statusinvest-crawler

replace (
	github.com/jasondavindev/statusinvest-crawler/greenblatt => ./greenblatt
	github.com/jasondavindev/statusinvest-crawler/requester => ./requester
)

go 1.15

require github.com/oleiade/reflections v1.0.1
