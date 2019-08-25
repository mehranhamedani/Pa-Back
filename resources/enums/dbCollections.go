package enums

// DBCollectionName type
type DBCollectionName string

// ToString func
func (dbCollectionName DBCollectionName) ToString() string {
	result := ""
	switch dbCollectionName {
	case DBCollectionNameUsers:
		result = "users"
	case DBCollectionNameOTPs:
		result = "otps"
	}
	return result
}

const (
	DBCollectionNameUsers DBCollectionName = "users"
	DBCollectionNameOTPs  DBCollectionName = "otps"
)
