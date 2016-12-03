package database

func GetUserToken(userId int) (string, error) {
	var token string

	q := "SELECT `refresh_token` FROM `oauth_token` WHERE `user_id` = ? LIMIT 1"
	err := DB.QueryRow(q, userId).Scan(&token)

	if err != nil {
		return "", err
	}

	return token, nil
}
