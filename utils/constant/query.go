package constant

const (
	USER_INSERT          = "INSERT INTO user_credential(id, username, password, role) VALUES ($1, $2, $3, $4);"
	USER_LIST            = "SELECT id, username, is_active, role FROM user_credential LIMIT $1 OFFSET $2;"
	USER_GET_TOTAL_ROWS  = "SELECT COUNT(*) FROM user_credential;"
	USER_GET             = "SELECT id, username, is_active, role FROM user_credential WHERE id=$1;"
	USER_UPDATE          = "UPDATE user_credential SET username=$1, role=$2, is_active=$3 WHERE id=$4;"
	USER_DELETE          = "DELETE FROM user_credential WHERE id=$1;"
	USER_GET_BY_USERNAME = "SELECT id, username, password, is_active, role FROM user_credential WHERE is_active=true and username=$1;"
	USER_UPDATE_PASSWORD = "UPDATE user_credential SET password=$1 WHERE id=$2;"

	PROFILE_INSERT = "INSERT INTO profile (id, name, address, phone, balance) VALUES ($1, $2, $3, $4, $5)"
	PROFILE_LIST   = "SELECT id, name, address, phone, balance FROM profile"
	PROFILE_GET    = "SELECT id, name, address, phone, balamce FROM profile WHERE id = $1"
	PROFILE_UPDATE = "UPDATE profile SET name = $1, address = $2, phone = $3 balance = $4 WHERE id = $5"
	PROFILE_DELETE = "DELETE FROM profile WHERE id = $1"

	MERCHANT_INSERT = "INSERT INTO profile (id, name_merchant_merchant, address, phone, balance) VALUES ($1, $2, $3, $4, $5)"
	MERCHANT_LIST   = "SELECT id, name_merchant, address, phone, balance FROM profile"
	MERCHANT_GET    = "SELECT id, name_merchant, address, phone, balamce FROM profile WHERE id = $1"
	MERCHANT_UPDATE = "UPDATE profile SET name_merchant = $1, address = $2, phone = $3 balance = $4 WHERE id = $5"
	MERCHANT_DELETE = "DELETE FROM profile WHERE id = $1"
	
	TRANSFER_INSERT = "INSERT INTO transfer (id, sender_id, receiver_id, amount, description) VALUES ($1, $2, $3, $4, $5);"
	TRANSFER_LIST   = "SELECT id, sender_id, receiver_id, amount, description FROM transfer LIMIT $1 OFFSET $2;"
	TRANSFER_GET    = "SELECT id, sender_id, receiver_id, amount, description FROM transfer WHERE id = $1;"
	TRANSFER_UPDATE = "UPDATE transfer SET sender_id = $2, receiver_id = $3, amount = $4, description = $5 WHERE id = $1;"
	TRANSFER_DELETE = "DELETE FROM transfer WHERE id = $1;"
	
	
)