package cbuild

import "fmt"

// TreeState represents a git tree state.
type TreeState int

// Enum values of TreeState.
const (
	TreeStateClean TreeState = iota
	TreeStateDirty
)

var (
	nameByTreeState = map[TreeState]string{
		TreeStateClean: "clean",
		TreeStateDirty: "dirty",
	}
)

func (s TreeState) String() string {
	return nameByTreeState[s]
}

// UnmarshalText implements encoding.TextUnmarshaler
func (s *TreeState) UnmarshalText(data []byte) error {
	switch string(data) {
	case "clean":
		*s = TreeStateClean
	case "dirty":
		*s = TreeStateDirty
	default:
		return fmt.Errorf("unknown state: %s", string(data))
	}
	return nil
}
