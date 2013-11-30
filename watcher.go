package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
	"regexp"
	"syscall"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event: ", ev)
				str := ev.String()
				log.Println(str)
				re, _ := regexp.Compile(`[^\"(.)\"$]+`)
				log.Println(re.FindStringSubmatch(str)[0])
				callShellParser()
			case err := <-watcher.Error:
				log.Println("error: ", err)
			}
		}
	}()

	err = watcher.Watch("data")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()

}

func callShellParser() {
	binary, lookErr := exec.LookPath("/var/www2/go/parser")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"pregame", "full"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
