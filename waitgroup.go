package main
import (
	"fmt"
	"sync"
)

func main(){
	fruits := []string{"apple","manggo","banana","rambutan"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruits(index, fruit, &wg)
	}

    wg.Wait()
}

func printFruits(index int, fruit string, wg *sync.WaitGroup){
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}

*yang sudah ada repo nya di github.

git add *
git status 
git commit -m "message"
git push origin master


project pertama da baru dipush pertama kali di github
git init
git add *
git status
git commit -m "message"
git remote add origin "url"
git push origin master
