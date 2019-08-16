package main

import (
    "fmt"
    "io"
    "os"
    "time"
)

func CopyFile(srcPath, dstPath string) (size int64, err error) {

    src, err := os.Open(srcPath)
    if err != nil {
        panic(err)
        // return
    }
    defer src.Close()

    dst, err := os.Create(dstPath)
    if err != nil {
        // panic(err)
        return
    }
    defer dst.Close()

    size, err = io.Copy(dst, src)

    return
}

func GetFileSize(filePath string) int64 {
    fileinfo, err := os.Stat(filePath)
    if  err != nil {
        panic(err)
    }
    return fileinfo.Size()
}

func main() {

    srcName := os.Args[1]
    dstName := os.Args[2]
    fmt.Printf("src: %s\n", srcName)
    fmt.Printf("dst: %s\n", dstName)

    srcFileSize := GetFileSize(srcName)
    fmt.Printf("src File Size: %d \n", srcFileSize)

    go func() {
        size, err := CopyFile(srcName, dstName)
        if  err != nil {
            panic(err)
        }
        fmt.Printf("Copied Size: %d\n", size)
    }()

    var dstFileSize int64
    for dstFileSize < srcFileSize {
        time.Sleep(1 * time.Second)
        dstFileSize = GetFileSize(dstName)
        fmt.Printf("dst File Size: %d \n", dstFileSize)
    }

    fmt.Println("END")
}
