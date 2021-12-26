package linkedliststack

import "gods/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Stack)(nil)
	var _ containers.JSONDeserializer = (*Stack)(nil)
}

func (stack *Stack) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

func (stack *Stack) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
