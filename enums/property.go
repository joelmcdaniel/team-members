package enums

// Property for member type property enum and key
type Property int

// Property enum ...
const (
	Duration Property = iota
	Role
)

// Name gets member type property name according to
// corresponding Property enum value (i.e. Type Key)
func (p Property) Name() string {
	return [...]string{"duration", "role"}[p]
}
