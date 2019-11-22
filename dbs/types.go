package dbs

type Note struct {
	Key       string `json:"key"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	UpdatedAt string `json:"update_time"`
	Content   string `json:"content"`
	Source    string `json:"source"`
	Editor    string `json:"Editor"`
	Files     string `json:"files"`
	Visit     int    `json:"visit"`
	Praise    int    `json:"praise"`
}

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar,omitempty"`
	Role   int    `json:"role"` // 0 管理员 1正常用户
	Editor string `json:"editor,omitempty"`
}

type PraiseLog struct {
	Key    string `json:"key"`
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
	Flag   bool   `json:"flag"`
}

type Message struct {
	Key     string `json:json:"key"`
	UserID  int    `json:"user_id"`
	NoteKey string `json:"note_key"`
	Content string `json:"content"`
	Praise  int    `json:"praise"`
}
