// Genre iota (Go's enum)

package dto

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
