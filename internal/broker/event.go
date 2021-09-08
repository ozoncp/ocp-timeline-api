package broker

type ActionType int

const (
	Create ActionType = iota
	Update
	Remove
)

type Message struct {
	Data map[string]interface{}
}

func (actionType ActionType) String() string {
	switch actionType {
	case Create:
		return "Created"
	case Update:
		return "Updated"
	case Remove:
		return "Removed"
	default:
		return "Unknown"
	}
}
