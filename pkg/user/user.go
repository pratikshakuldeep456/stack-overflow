package user

type User struct {
	ID              int
	Name            string
	ReputationScore int
	//Posts           []*post.Question
	//mutex for thread safety
}
