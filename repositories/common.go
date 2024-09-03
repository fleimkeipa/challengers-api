package repositories

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getPaginationOpts(l, s uint) options.FindOptions {
	var limit = int64(l)
	if limit == 0 {
		limit = 30
	}
	var skip = int64(s)
	return options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
}
