package future

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func timeout(t *testing.T, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	t.Log("Timeout!")
	t.Fail()
	wg.Done()
}

func TestStringOrErrorExecute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(e error) {
			t.Fail()
		})

		future.Execute(func() (string, error) {
			return "Hello world!", nil
		})

		go timeout(t, &wg)

		wg.Wait()

	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)

		future.Success(func(s string) {
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})

		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
		wg.Wait()

	})
}
