package main

import (
	"fmt"
	"log"
	"os"

	"github.com/isolk/go-lsp/logs"
)

func test() {
	f, e := os.Create("log.txt")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	logs.Init(log.New(f, "", log.LstdFlags))

	// res, err := ast.SemanticToken("a.sdp")
	// fmt.Printf("res: %v, err: %v\n", res, err)
}

func main() {
	// ast.TestFormatCode()
	startLspServer()
	// a := []string{"IDNum", "ID", "IDCard", "IDCardNum", "IDCardNumber"}
	// b := "numcard"
	// res := find(a, b)
	// fmt.Println(res)

	// fA := "/a/b/c/d/e.go"
	// fB := "/a/b/c/d/f.go"
	// fmt.Println(findRelativeFileName(fA, fB))
}

// func findRelativeFileName(targetFileName, newFileName string) string {
// 	// 使用 filepath.Dir() 获取文件路径的父级目录
// 	dirA := filepath.Dir(targetFileName)
// 	dirB := filepath.Dir(newFileName)

// 	// 使用 filepath.Rel() 获取相对路径
// 	relativePath, _ := filepath.Rel(dirA, dirB)

// 	// 处理相对路径中的 "." 和 ".."
// 	// relativePath = filepath.Clean(relativePath)

// 	return relativePath + "/" + filepath.Base(newFileName)
// }
