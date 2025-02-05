package post

type Question struct {
	ID        int
	UserID    int
	Title     string
	Content   string
	Tag       []string
	Comments  []*Comment
	Answers   []*Answer
	VoteCount int
}
