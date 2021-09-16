package captures_mongo

type Upsert struct {
	Set         interface{} `bson:"$set,omitempty"`
	SetOnInsert interface{} `bson:"$setOnInsert,omitempty"`
}

func SetStatusOnInsert(status string) SetOnInsert {
	return SetOnInsert{Status: status}
}

type SetOnInsert struct {
	Status string `bson:"status"`
}
