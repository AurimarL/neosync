package neosync_benthos

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func BuildBenthosTable(schema, table string) string {
	if schema != "" {
		return fmt.Sprintf("%s.%s", schema, table)
	}
	return table
}

func HashBenthosCacheKey(jobId, runId, table, col string) string {
	return ToSha256(fmt.Sprintf("%s.%s.%s.%s", jobId, runId, table, col))
}

func ToSha256(input string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(input)))
}

// checks if the error message is critical
func IsCriticalError(errMsg string) bool {
	// list of known error messages for when max connections are reached
	criticalErrors := []string{
		"violates foreign key constraint",
		"duplicate key value violates unique constraint",
		"duplicate entry",
		"cannot add or update a child row",
		"a foreign key constraint fails",
		"could not identify an equality operator",
		"violates not-null constraint",
		"failed to send message to redis_hash_output",
		"mapping returned invalid key type",
		"invalid input syntax",
		"incorrect datetime value",
		"incorrect date value",
		"incorrect time value",
		"does not exist",
		"syntax error at or near",
		"ON CONFLICT DO UPDATE requires inference specification or constraint name",
		"transaction has already been committed or rolled back",
		"missing redis client",
		"violates check constraint",
		"ON CONFLICT does not support deferrable unique constraints",
		"ON CONFLICT",
		"SQLSTATE", // any sqlstate error should result in ending
		"goqu_encode_error",
		"doesn't have a default value",
		"column does not allow nulls",
	}

	for _, errStr := range criticalErrors {
		if containsIgnoreCase(errMsg, errStr) {
			return true
		}
	}
	return false
}

// checks if the error message is critical for the generate job
func IsGenerateJobCriticalError(errMsg string) bool {
	criticalErrors := []string{
		"violates foreign key constraint",
		"cannot add or update a child row",
		"a foreign key constraint fails",
		"could not identify an equality operator",
		"violates not-null constraint",
		"invalid input syntax",
		"incorrect datetime value",
		"incorrect date value",
		"incorrect time value",
		"does not exist",
		"syntax error at or near",
		"doesn't have a default value",
		"column does not allow nulls",
	}

	for _, errStr := range criticalErrors {
		if containsIgnoreCase(errMsg, errStr) {
			return true
		}
	}
	return false
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func IsForeignKeyViolationError(errMsg string) bool {
	foreignKeyViolationErrors := []string{
		"violates foreign key constraint",
		"a foreign key constraint fails",
		"insert statement conflicted with the foreign key constraint",
	}

	for _, errStr := range foreignKeyViolationErrors {
		if containsIgnoreCase(errMsg, errStr) {
			return true
		}
	}
	return false
}

func ShouldRetryInsert(errMsg string, shouldCheckForForeignKeyViolation bool) bool {
	if shouldCheckForForeignKeyViolation && IsForeignKeyViolationError(errMsg) {
		return true
	}
	otherErrors := []string{
		"ON CONFLICT DO UPDATE command cannot affect row a second time",
	}
	for _, errStr := range otherErrors {
		if containsIgnoreCase(errMsg, errStr) {
			return true
		}
	}
	return false
}
