package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	var comslice []string

	//os.Argsの0番目以降をスライスに
	comslice = os.Args[1:]

	if len(comslice) == 0 {
		fmt.Println("実行するコマンドを指定してください。")
		fmt.Printf(" 例）%s コマンド コマンドライン引数1 コマンドライン引数2... \n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	cmd := exec.Command(comslice[0], comslice[1:]...)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	exitCode := cmd.ProcessState.ExitCode() //終了コード取得

	if err != nil {
		fmt.Print("\n[Error] ")
		os.Stderr.WriteString(err.Error())
		for g, h := range comslice[:] {
			if g == 0 {
				fmt.Printf("\nコマンド:%s\n", h)
			} else {
				fmt.Printf("引数%d:%s\n", g, h)
			}

		}

	}

	fmt.Print(string(cmdOutput.Bytes()))
	fmt.Println("\n終了コード:", exitCode)

}
