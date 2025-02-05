package post

type Answer struct {
	ID         int
	UserID     int
	QuestionID int
	Content    string
	VoteCount  int
	Comment    []*Comment
}
