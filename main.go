package main

import (
	"os/exec"
	"os/user"
	"fmt"
	"os"
	"encoding/gob"
	"log"
	//"strings"
	"syscall"
	"os/signal"
)


func main() {
	checkDockerStatus()
	createOneoffContainer()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	go func() {
		s := <-sigc

	}()

	//for i := 0; i < len(os.Args); i++ {
	//	switch strings.ToLower(os.Args[i]) {
	//
	//		// TODO reload
	//		break
	//	case "oneoff":
	//
	//		// TODO oneoff
	//		break
	//
	//	}
	//}
}

func reloadContainer() {

}

func createOneoffContainer() (stdOutput string, err error) {
	stdOutput, err = executeCommandSilent("docker", "run", "-td", "--rm", "technowizard/tcnj_linux_container:latest")
	return
}

// checks if docker is in a runnable state
func checkDockerStatus() {
	cmd := exec.Command("docker", "version")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error: Docker is not installed or not running. '%v'\nFatal. Terminating\n", err)
	}
}

func writeGob(filePath string,object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string,object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// gets the executable data file
func getDataFile() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(usr.HomeDir)
}

func containerExists() {

}

func enableScreenBuffer() {

}

func disableScreenBuffer() {

}


func executeCommandSilent(command string, args ...string) (output string, err error) {
	var buffer []byte
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	buffer, err = cmd.Output()
	output = string(buffer)
	return
}

