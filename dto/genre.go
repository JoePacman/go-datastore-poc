// Genre iota (Go's enum)

package dto

import (
	"bytes"
	"encoding/json"
)

type Genre int

const (
	Action Genre = iota
	Romance
	SciFi
	Drama
)

func (g Genre) String() string {
	return [...]string{"Action", "Romance", "SciFi", "Drama"}[g]
}

var toID = map[string]Genre{
	"Action":  Action,
	"Romance": Romance,
	"SciFi":   SciFi,
	"Drama":   Drama,
}

var toString = map[Genre]string{
	Action:  "Action",
	Romance: "Romance",
	SciFi:   "SciFi",
	Drama:   "Drama",
}

func (g Genre) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[g])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (g *Genre) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*g = toID[j]
	return nil
}
