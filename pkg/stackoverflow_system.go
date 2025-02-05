package pkg

import (
	"errors"
	"pratikshakuldeep456/stack-overflow/pkg/post"
	"pratikshakuldeep456/stack-overflow/pkg/user"
	"strings"
)

type StackOverflow struct {
	Users     map[int]*user.User
	Questions map[int]*post.Question
	Answers   map[int]*post.Answer
	Tags      map[string][]*post.Question
}

func NewStackOverflowSystem() *StackOverflow {
	return &StackOverflow{
		Users:     make(map[int]*user.User),
		Questions: make(map[int]*post.Question),
		Answers:   make(map[int]*post.Answer),
		Tags:      make(map[string][]*post.Question)}
}

func (so *StackOverflow) NewUser(name string) *user.User {
	user := &user.User{ID: user.GenerateID(), Name: name, ReputationScore: 0}

	so.Users[user.ID] = user

	return user

}
func (so *StackOverflow) PostQues(userID int, title, content string, tags []string) (*post.Question, error) {
	//check user if exists or not

	if _, exists := so.Users[userID]; !exists {
		return nil, errors.New("user does not exist in system")
	}
	ques := &post.Question{
		ID:        user.GenerateID(),
		UserID:    userID,
		Title:     title,
		Content:   content,
		Tag:       tags,
		Comments:  []*post.Comment{},
		Answers:   []*post.Answer{},
		VoteCount: 0,
	}

	so.Questions[ques.ID] = ques

	for _, tag := range tags {
		so.Tags[tag] = append(so.Tags[tag], ques)
	}

	return ques, nil
}

func (so *StackOverflow) PostAnswer(userID, postID int, content string) (*post.Answer, error) {

	if _, exists := so.Users[userID]; !exists {
		return nil, errors.New("user does not exist in system")
	}
	question, exists := so.Questions[postID]
	if !exists {
		return nil, errors.New("question doesnt found")
	}

	ans := &post.Answer{
		ID:         user.GenerateID(),
		UserID:     userID,
		QuestionID: postID,
		Content:    content,
		VoteCount:  0,
		Comment:    []*post.Comment{},
	}
	so.Answers[ans.ID] = ans

	question.Answers = append(question.Answers, ans)
	return ans, nil
}

func (so *StackOverflow) PostComment(userID, contentID int, isQuestion bool, content string) (*post.Comment, error) {
	if _, exists := so.Users[userID]; !exists {
		return nil, errors.New("user does not exist in system")
	}

	comment := &post.Comment{
		UserID:  userID,
		Content: content,
	}

	if isQuestion {
		question, exists := so.Questions[contentID]
		if !exists {
			return nil, errors.New("question not found")
		}
		question.Comments = append(question.Comments, comment)
	} else {
		ans, exists := so.Answers[contentID]
		if !exists {
			return nil, errors.New("answer not found")
		}
		ans.Comment = append(ans.Comment, comment)

	}
	return comment, nil

}

func (so *StackOverflow) Vote(userID, contentID int, isQuestion, upvote bool) error {
	var ownerID int
	if isQuestion {
		q, exists := so.Questions[contentID]
		if !exists {
			return errors.New("question not found")
		}
		ownerID = q.UserID
		if upvote {
			q.VoteCount++
		} else {
			q.VoteCount--
		}
	} else {
		a, exists := so.Answers[contentID]
		if !exists {
			return errors.New("answer not found")
		}
		ownerID = a.UserID
		if upvote {
			a.VoteCount++
		} else {
			a.VoteCount--
		}

	}
	//update reputation
	u, exists := so.Users[ownerID]
	if exists {
		if upvote {
			u.ReputationScore += 10
		} else {
			u.ReputationScore -= 2
		}
	}
	return nil
}
func (so *StackOverflow) SearchQuestion(keyword string, tag string, userID int) ([]*post.Question, error) {
	var ans []*post.Question

	if tag != "" {
		if i, exists := so.Tags[tag]; exists {
			ans = append(ans, i...) //explain??
		}
	}
	for _, val := range so.Questions {
		if val.UserID == userID {
			ans = append(ans, val)
		} else if keyword != "" && (strings.Contains(strings.ToLower(val.Title), strings.ToLower(keyword))) || (strings.Contains(strings.ToLower(val.Content), strings.ToLower(keyword))) {
			ans = append(ans, val)
		}
	}

	return ans, nil
}
