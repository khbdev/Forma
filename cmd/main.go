package main

import (
	"forma/internal/cache"
	"forma/internal/config"
	repository "forma/internal/repostory"
	"forma/internal/service"
	loadenv "forma/pkg/loadEnv"
	"log"
)




func main(){
	loadenv.LoadEnv()

	postgres, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	redis, err := config.NewRedisClient()
		if err != nil {
		log.Fatal("Error: ", err)
	}


	


	cache := cache.NewLeadCache(redis)

   repos := repository.NewLeadRepository(postgres)

   srv := service.NewLeadService(repos)



}