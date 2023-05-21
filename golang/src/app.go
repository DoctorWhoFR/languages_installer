package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func execBasicCmds(dir, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	out, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return out, nil
}

const LOCAL_FOLDER string = "/usr/local"

func main() {

	if len(os.Args) == 1 {
		panic("should put an arg")
	}

	arg1 := (os.Args[1]) // should be golang tar link file

	dl_file := "https://go.dev/dl/"

	// 1 - Should remove current Golang Folder
	_, err := execBasicCmds(LOCAL_FOLDER, "rm", "-R", "gotest")

	if err != nil {
		log.Println("file does not exist")
	}

	log.Println("old golang folder removed.")

	// 2 - should download golang tar file
	_, err2 := execBasicCmds(LOCAL_FOLDER, "wget", dl_file+arg1)

	if err2 != nil {
		panic(err2)
	}

	log.Println("Check if file is downloaded")

	ls_out, err2_1 := execBasicCmds(LOCAL_FOLDER, "ls")

	if err2_1 != nil {
		panic(err2_1)
	}

	if strings.Contains(string(ls_out), arg1) {
		log.Println("we have the golang file, untaring it")
	}

	// 3 - untar the file
	_, err3 := execBasicCmds(LOCAL_FOLDER, "tar", "-C", "/usr/local", "-xvf", arg1)

	if err3 != nil {
		panic(err3)
	}

	// 4 - check go folder

	log.Println("Check if folder is present")

	ls_out4, err4 := execBasicCmds(LOCAL_FOLDER, "ls")

	if err4 != nil {
		panic(err4)
	}

	if strings.Contains(string(ls_out4), "go") {
		log.Println("we have the golang folder")
	}

	// finished basic install

	// 5 - Send path command
	log.Println("Everything is ready to be used.\nPLEASE, you should execute this 2 steps before using golang :")
	log.Println("1 - Execute command:", "'", "declare", "-x", "PATH=$PATH:/usr/local/go/bin", "'")
	log.Println("2 - Execute command:", "'", "go", "version", "'")

}
