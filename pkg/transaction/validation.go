package transaction

import (
	"fmt"
	"sync"
)

type Validation interface {
	Valid(trx Transaction, wg *sync.WaitGroup, result chan<- bool)
}

type MissingAmount struct {
}

type MissingAccountID struct {
}

var _ Validation = (*MissingAmount)(nil)
var _ Validation = (*MissingAccountID)(nil)

func Validate(trx Transaction) {
	var wg sync.WaitGroup

	result := make(chan bool)
	validations := make([]Validation, 0)

	validations = append(validations, &MissingAmount{})
	validations = append(validations, &MissingAccountID{})

	for _, validator := range validations {
		wg.Add(1)
		go validator.Valid(trx, &wg, result)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	var results []bool

	for r := range result {
		results = append(results, r)
	}

	fmt.Println(results)
}

func (m *MissingAmount) Valid(trx Transaction, wg *sync.WaitGroup, result chan<- bool) {
	result <- false

	defer wg.Done()

}

func (m *MissingAccountID) Valid(trx Transaction, wg *sync.WaitGroup, result chan<- bool) {

	result <- true

	defer wg.Done()

}
