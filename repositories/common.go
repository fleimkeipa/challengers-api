package repositories

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
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

func getFilterValues(key, value string) bson.M {
	var splitted = strings.Split(value, ",")
	if len(splitted) > 1 {
		return bson.M{
			key: bson.M{
				"$in": splitted,
			},
		}
	}

	return bson.M{key: value}
}
