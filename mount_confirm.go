package main

import (
	// 標準ライブラリ
	"flag"
	"fmt"
	"os"

	// 外部ライブラリ
	"github.com/docker/docker/pkg/mount"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(os.Args) < 2 {
		// 対象ファイル名が引数にない場合は標準入力から読み込み
		fmt.Println("Usage: ", os.Args[0]+" mount-point")
		return
	}

	// 引数のマウントポイントを表示
	fmt.Println("mount-point: ", args[0])

	// マウントポイントのマウント状態を確認
	// Mounted determines if a specified mountpoint has been mounted. On Linux it looks at /proc/self/mountinfo.
	bSt, err := mount.Mounted(args[0])

	// マウント状態確認結果を表示
	fmt.Println("mount (true/false): ", bSt)
	fmt.Println("mount (error): ", err)
}
