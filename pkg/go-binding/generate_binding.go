package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

const (
	buildDir  = "../../builds/"
	targetDir = "../ethereum/"
)

func main() {
	//contractNames := []string{"CustodianWalletProxy"}
	contractNames := []string{"CustodianWalletLogic", "Factory", "Escrow", "USDC"}

	cmd := exec.Command("which","abigen")
	err := cmd.Run()

	if err != nil {
		log.Fatal("ethereum-abigen is not a command: ",err)
	}

	for _, contractName := range contractNames {
		buildFilename := buildDir + contractName
		bindingFilename := targetDir + contractName + ".go"
		cmd = exec.Command(
			"abigen", fmt.Sprintf("--bin=%s.bin", buildFilename),
			fmt.Sprintf("--abi=%s.abi", buildFilename),
			fmt.Sprintf("--pkg=%s", contractName),
			fmt.Sprintf("--out=%s", bindingFilename),
		)

		err := cmd.Start()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Running... %s", cmd)
		err =  cmd.Wait()
		if err != nil {
			log.Fatalln(err)
		}
		bindFile, err := ioutil.ReadFile(bindingFilename)
		if err != nil {
			log.Fatalln("Cannot open: ",bindingFilename, err)
		}

		fileOutput := strings.Replace(string(bindFile), "package " + contractName, "package ethereum", 1)
		err = ioutil.WriteFile(bindingFilename, []byte(fileOutput), 0644)
		if err != nil {
			log.Fatalln("Cannot write to: ",bindingFilename,err)
		}
	}
}
