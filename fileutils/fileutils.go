package fileutils

import (
	"log"
	"os"
	"time"
)

func DeleteFolder(folder string) error {

	errchan := make(chan error)

	go func() {
		time.Sleep(5 * time.Second)
		err := os.RemoveAll(folder)
		errchan <- err
	}()

	select {
	case err := <-errchan:
		return err
	case <-time.After(30 * time.Second):
		log.Println("Did not resolve in time")
	}

	return nil
}
