package main

import (
	"fmt"
	"pratikshakuldeep456/stack-overflow/pkg"
)

func main() {

	so := pkg.NewStackOverflowSystem()

	user1 := so.NewUser("Rahul")
	fmt.Println("creating a user", *user1)
	user2 := so.NewUser("Rolly")
	fmt.Println("creating a user", *user2)

	q1, _ := so.PostQues(user1.ID, "How does Go handle concurrency?", "Goroutines and channels...", []string{"golang", "concurrency"})
	q2, _ := so.PostQues(user2.ID, "What is a mutex in Go?", "A mutex ensures mutual exclusion...", []string{"golang", "mutex"})

	so.PostAnswer(user1.ID, q2.ID, "Go uses Goroutines and channels.")
	so.PostAnswer(user1.ID, q2.ID, "A mutex ensures mutual exclusion.")

	so.PostAnswer(q2.UserID, q1.ID, " A mutex ensures mutual exclusion.")
	so.PostComment(user1.ID, q1.UserID, true, "understood")

	so.Vote(user1.ID, q1.ID, true, true)

	results, _ := so.SearchQuestion("mutex", "golang", user1.ID)
	fmt.Println("Questions with 'golang' tag:", len(results))

	fmt.Println("Questions with 'golang' tag:", len(results))

}
