package utils

type ISerializable interface {
	Serialize() map[string]interface{}
}

type Serializer struct {}

func (Serializer) Serialize(serializables []ISerializable) []map[string]interface{} {
	var collection []map[string]interface{}
	for _, serializable := range serializables {
		collection = append(collection, serializable.Serialize())
	}
	return collection
}
