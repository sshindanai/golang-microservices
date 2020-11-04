package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sshindanai/golang-microservices/src/api/api/errors"
	"github.com/sshindanai/golang-microservices/src/api/domain/repositories"
	"github.com/sshindanai/golang-microservices/src/api/services"
)

var (
	success map[string]string
	failed  map[string]string
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequest() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/home/sshindanai/Documents/workspace/go/src/github.com/sshindanai/golang-microservices/concurrency/requests.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scan line by line til the end of the file
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	return result
}

func main() {
	requests := getRequest()

	fmt.Println(fmt.Sprintf("about to process %d requests", len(requests)))

	input := make(chan createRepoResult)
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(input, &wg)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(buffer, input, request)
	}

	wg.Done()

	wg.Wait()
	close(input)
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error.Message()
			continue
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}

	// Release the slots from capacity
	<-buffer
}
