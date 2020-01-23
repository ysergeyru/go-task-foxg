package pg

import (
	"database/sql"
)

// GetConnectionsByUserIDs returns all users connection log entries by user IDs
func GetConnectionsByUserIDs(firstUserID, secondUserID int) ([]string, error) {
	userConnections := []string{}
	if err := DB().Select(&userConnections, "SELECT ip_addr FROM conn_log WHERE user_id in ($1, $2)", firstUserID, secondUserID); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return userConnections, nil
}
