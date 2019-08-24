package enums

// DBCollectionName type
type DBCollectionName string

// ToString func
func (dbCollectionName DBCollectionName) ToString() string {
	result := ""
	switch dbCollectionName {
	case DBCollectionNameUsers:
		result = "users"
	}
	return result
}

const (
	// DBCollectionNameUsers users
	DBCollectionNameUsers DBCollectionName = "users"
)
