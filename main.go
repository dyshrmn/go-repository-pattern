package main

import (
	"fmt"

	"github.com/dyshrmn/go-repository-pattern/entity"
	"github.com/dyshrmn/go-repository-pattern/repository"
)

func main() {
	userRepository := repository.UserRepository{repository.InitBase()}
	postRepository := repository.PostRepository{repository.InitBase()}

	userOne := entity.User{
		Name:     "dyshrmn-",
		Username: "dyshrmn",
		Password: "dyshrmn?",
	}
	userTwo := entity.User{
		Name:     "kyujay-",
		Username: "kyujay",
		Password: "kyujay?",
	}

	userRepository.Create(userOne)
	userRepository.Create(userTwo)

	postRepository.Create(entity.Post{
		Name: "Part 1: How to ...",
		User: userOne,
	})
	postRepository.Create(entity.Post{
		Name: "Part 2: How to ...",
		User: userTwo,
	})

	_users := userRepository.FindAll()
	for _id, _user := range _users {
		user := _user.(entity.User)
		fmt.Printf("user> id: %d name: %s username: %s password: %s\n",
			_id, user.Name, user.Username, user.Password)
	}

	_posts := postRepository.FindAll()
	for _id, _post := range _posts {
		post := _post.(entity.Post)
		fmt.Printf("post> id: %d user: %s name: %s\n", _id, post.User.Name, post.Name)
	}
}
