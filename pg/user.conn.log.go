package pg

import (
	"database/sql"
)

// GetConnectionsByUserIDs returns all users connection log entries by user IDs
func GetConnectionsByUserIDs(firstUserID, secondUserID int) ([]string, error) {
	userIPs := []string{}

	if err := DB().Select(&userIPs, "SELECT ip_addr FROM user_id_to_ip WHERE user_id in ($1, $2)", firstUserID, secondUserID); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return userIPs, nil
}
