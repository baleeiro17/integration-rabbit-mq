package json

import (
	"bytes"
	"encoding/json"
	"integration-rabbit-mq/internal/models"
)

func Serialize(msg models.Message) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func Deserialize(b []byte) (models.Message, error) {
	var msg models.Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
