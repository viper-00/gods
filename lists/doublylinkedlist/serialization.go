package doublylinkedlist

import (
	"encoding/json"
	"gods/containers"
)

func assertSerializatioImplentation() {
	var _ containers.JSONSerializer = (*List)(nil)
	var _ containers.JSONDeserializer = (*List)(nil)
}

func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

func (list *List) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err != nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
