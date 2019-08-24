package enums

type DBCollectionName string

// ToString func
func (dbCollectionName DBCollectionName) ToString() string {
	result := ""
	switch dbCollectionName {
	case DBCollectionName_Users:
		result = "users"
	}
	return result
}

const (
	DBCollectionName_Users DBCollectionName = "users"
)
