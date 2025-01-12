package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/qwerqy/social-api-go/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "david", "emma", "frank", "grace", "harry", "ivy", "jack",
	"kate", "leo", "mia", "noah", "oliver", "paul", "quincy", "rose", "sophie", "thomas",
	"ursula", "victor", "william", "xavier", "yara", "zoe", "adam", "bella", "clara", "daniel",
	"edward", "fiona", "george", "hannah", "isabella", "james", "kevin", "lucas", "mason", "nina",
	"oscar", "peter", "quinn", "rachel", "samuel", "tina", "ursula", "vincent", "wesley", "xander",
}

var titles = []string{
	"Mastering Time Management: Tips for Busy Professionals",
	"The Ultimate Guide to Boosting Your Productivity in 2024",
	"How to Stay Motivated When Working on Long-Term Goals",
	"10 Essential Tools for Remote Work Success",
	"The Art of Simplifying Complex Projects: A Step-by-Step Approach",
	"Overcoming Procrastination: Strategies That Actually Work",
	"Embracing Change: How to Thrive in a Dynamic Environment",
	"Top 5 Coding Practices Every Developer Should Know",
	"Building Scalable Web Apps with Go: A Beginner’s Guide",
	"Why Every Engineer Should Learn the Go Programming Language",
	"Debugging Go Applications: Tips and Techniques",
	"Best Practices for Writing Clean and Efficient Code in Go",
	"The Rise of Go: Why It’s the Future of Software Development",
	"How to Write Performant APIs with Go",
	"Error Handling in Go: Patterns and Best Practices",
	"Exploring Concurrency in Go: Channels, Goroutines, and More",
	"From Novice to Pro: A Developer's Journey in Learning Go",
	"Testing in Go: Ensuring Reliability and Stability in Your Code",
	"An Introduction to Go Modules and Dependency Management",
	"How to Build a Simple Web Server in Go",
}

var contents = []string{
	"Learn how to manage your time effectively and get more done in less time.",
	"Discover actionable tips to boost your productivity and achieve your goals.",
	"Stay motivated with these proven strategies for tackling long-term projects.",
	"Explore the top tools that make remote work seamless and efficient.",
	"Break down complex projects into manageable steps for better results.",
	"Overcome procrastination with these simple yet powerful techniques.",
	"Adapt to change and thrive in today’s fast-paced work environments.",
	"Improve your coding skills with these essential programming practices.",
	"Get started with building scalable web applications using Go.",
	"Learn why Go is a must-know language for modern developers.",
	"Master debugging techniques to quickly identify and fix issues in Go.",
	"Write clean and efficient Go code with these best practices.",
	"Explore why Go is the language of choice for high-performance applications.",
	"Learn how to build fast and reliable APIs using Go’s features.",
	"Understand error handling in Go with practical patterns and examples.",
	"Dive into Go’s concurrency model with channels and goroutines.",
	"Follow a developer's journey to becoming a Go expert.",
	"Ensure code quality and reliability with testing best practices in Go.",
	"Manage dependencies effortlessly with Go modules.",
	"Build a simple yet powerful web server with Go in just a few steps.",
}

var tags = []string{
	"go",
	"programming",
	"web development",
	"api design",
	"productivity",
	"software engineering",
	"golang",
	"debugging",
	"clean code",
	"best practices",
	"concurrency",
	"remote work",
	"coding tips",
	"error handling",
	"developer tools",
	"learning go",
	"scalable systems",
	"modern development",
	"performance",
	"code quality",
}

var comments = []string{
	"Great article! Very informative and easy to follow.",
	"I had no idea about this feature in Go. Thanks for sharing!",
	"Could you provide more examples for better understanding?",
	"This really helped me solve a problem I was stuck on. Thanks!",
	"Looking forward to more content like this.",
	"Can you explain the performance impact of this approach?",
	"Fantastic write-up! It clarified many doubts I had.",
	"I think there’s a typo in the code snippet. Can you double-check?",
	"Thank you for breaking down a complex topic so well.",
	"This was exactly what I needed for my project!",
	"Can you recommend more resources for learning Go?",
	"Interesting perspective. I hadn’t thought about it this way before.",
	"Great post! Would love to see more on concurrency in Go.",
	"How does this compare to other programming languages?",
	"This saved me hours of debugging. Appreciate it!",
	"Do you have any advice for beginners in Go?",
	"I tried this approach, and it worked like a charm.",
	"Thanks for sharing this! It’s a game changer for me.",
	"Could you dive deeper into testing best practices in Go?",
	"Awesome post! The examples were really helpful.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()

			log.Println("Error seeding user: ", err)

			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error seeding post: ", err)

			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.CreateByPostID(ctx, comment); err != nil {
			log.Println("Error seeding post: ", err)

			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		cms[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
