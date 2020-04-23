//Issues prints a table of github issues
package main

import (
	"fmt"
	"github"
	"log"
	"os"
	"strconv"
)

func main() {
	date, err2 := strconv.Atoi(os.Args[1])
	if err2 != nil {
		log.Fatal(err2)
	}
	println(date)
	result, err := github.SearchIssues(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.Year() == date {
			fmt.Println(
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}
}
