package main

import (
	//	"fmt"
	"goproject_SBG-bot/api"
	"goproject_SBG-bot/repository"
	"goproject_SBG-bot/service"
)

func main() {

	repo := repository.New()
	srv := service.New(repo)
	api.Run(srv)

	//	fmt.Println(repo.Persons_name["Bora"])
	//	fmt.Println(srv.Repo)
	//+++++++++++++++++++++++++++++++++++++++++++
	//	text_vvod := "/list_name"
	//	replyMsg := srv.Distribution_answers(text_vvod, 3746)
	//	fmt.Println(replyMsg)
	//	fmt.Println(srv.Repo)
	//	fmt.Println(srv.Get("Bora"))
}
