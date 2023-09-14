package entity

type Photo struct {
	Id        string `db:"id"`
	Title     string `db:"title"`
	Caption   string `db:"caption"`
	Url       string `db:"url"`
	UserId    string `db:"user_id"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}
