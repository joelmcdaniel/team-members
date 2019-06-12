package enums

// Type for member type enum and key
type Type int

// Type enum
const (
	Contractor Type = iota
	Employee
	end
)

// Value gets member type value according to Type
// enum value (i.e. Type Key)
func (t Type) Value() string {
	return [...]string{"Contractor", "Employee"}[t]
}

// IsValidTypeKeyValue checks if value of member
// type key value is valid
func IsValidTypeKeyValue(value int) bool {
	return value < int(end)
}
