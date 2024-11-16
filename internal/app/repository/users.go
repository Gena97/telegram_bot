package repository

import (
	"fmt"

	"github.com/Gena97/telegram_bot/internal/app/model"
)

func (r *PGXRepository) GetUserInfosMap() (map[int64]model.User, error) {
	rows, err := r.mainDB.Query(`SELECT user_id, COALESCE(username, ''), first_name, COALESCE(last_name, ''), is_admin FROM users`)
	if err != nil {
		return nil, fmt.Errorf("error querying users from database: %w", err)
	}
	defer rows.Close()

	userMap := make(map[int64]model.User)

	for rows.Next() {
		var userInfo model.User

		if err := rows.Scan(&userInfo.UserID, &userInfo.Username, &userInfo.FirstName, &userInfo.LastName, &userInfo.IsAdmin); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		userMap[userInfo.UserID] = userInfo
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return userMap, nil
}
