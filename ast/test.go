package ast

import (
	"fmt"
	"os"
)

func TestIncludeManager() {
	in := NewIncludeManager()
	in.addRef("a", "b")
	res := in.recaculateRealRef()
	fmt.Println(res)

	in.addRef("a", "c")
	in.addRef("d", "c")
	res = in.recaculateRealRef()
	fmt.Println(res)

	in.addRef("c", "e")
	res = in.recaculateRealRef()
	fmt.Println(res)
}

func TestFormatCode() {
	fileName := "test_sdp_files/test_format.sdp"
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile(fileName, string(content))

	res, err := FormatCode(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)

	OnFileChanged(fileName, res)

	res, err = FormatCode(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)

	// ers, err := CheckSyntax(fileName)
	// fmt.Println(ers, err)
}

func TestCheckSyntax() {
	content, err := os.ReadFile("test_sdp_files/semantic_error.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("a.sdp", string(content))

	res, err := CheckSyntax("a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v.Message)
	}
}

func TestSemanticToken() {
	content, err := os.ReadFile("a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("a.sdp", string(content))

	res, err := SemanticToken("a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func TestHover() {
	content, err := os.ReadFile("a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("a.sdp", string(content))

	res, err := Hover("a.sdp", 9, 19)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

func TestDefinition() {
	content, err := os.ReadFile("a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("a.sdp", string(content))

	res, err := Definition("a.sdp", 9, 19)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res == nil {
		fmt.Println("nil")
		return
	}

	for _, v := range *res {
		fmt.Printf("uri: %s, range: %v\n", v.TargetUri, v.TargetRange)
	}

}

func TestInclude() {
	content, err := os.ReadFile("test_sdp_files/include_a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("test_sdp_files/include_a.sdp", string(content))

	res, err := CheckSyntax("test_sdp_files/include_a.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v.Message)
	}
}

func TestFileChanged() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		// 删除tmp目录
		os.RemoveAll("tmp")
	}()
	// 创建tmp目录
	// 创建1.sdp文件，2.sdp文件，1.sdp文件内容为#include "2.sdp"，并且用到了2.sdp中的类型
	// 打印错误
	// 提示输入enter
	// 修改2.sdp文件内容，将2.sdp文件中的类型改为其他类型
	// 打印错误

	// 创建tmp目录
	os.Mkdir("tmp", os.ModePerm)
	// 创建1.sdp文件
	file1Content := `#include "2.sdp"
module MTTD
{
    struct A
    {
        0 optional B a;
    };
};
	`
	os.WriteFile("tmp/1.sdp", []byte(file1Content), os.ModePerm)

	// 创建2.sdp文件
	file2Content := `module MTTD
{
	struct B{};
};
	`
	os.WriteFile("tmp/2.sdp", []byte(file2Content), os.ModePerm)

	// 打印错误
	content, err := os.ReadFile("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("tmp/1.sdp", string(content))

	res, err := CheckSyntax("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v.Message)
	}

	fmt.Println("file changed!!!")

	// 修改2.sdp文件内容
	file2Content = `module MTTD
{
	struct C{};
};
	`
	os.WriteFile("tmp/2.sdp", []byte(file2Content), os.ModePerm)
	OnFileChanged("tmp/2.sdp", file2Content)

	res, err = CheckSyntax("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range res {
		fmt.Println(v.Message)
	}
}

func TestSError() {
	content, err := os.ReadFile("test_sdp_files/test_error.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("test_sdp_files/test_error.sdp", string(content))

	res, err := CheckSyntax("test_sdp_files/test_error.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range res {
		fmt.Println(v.Message)
	}
}

// 检查先前的错误是否被修复，假设1.sdp忘了引用2.sdp，现在引用了，应该没有错误
func TestIncludeError() {
	os.Mkdir("tmp", os.ModePerm)
	// 创建1.sdp文件
	file1Content := `
	#include"2.sdp"
	module MTTD
{
	struct A
	{
		0 optional B a;
	};
};
	`
	os.WriteFile("tmp/1.sdp", []byte(file1Content), os.ModePerm)

	// 创建2.sdp文件
	file2Content := `module MTTD
{
	struct B{};
};
	`
	os.WriteFile("tmp/2.sdp", []byte(file2Content), os.ModePerm)

	// 打印错误
	OnOpenFile("tmp/1.sdp", string(file1Content))
	res, err := CheckSyntax("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range res {
		fmt.Println(v.Message)
	}

	fmt.Println("file changed1!!!")

	// 修改1.sdp文件内容
	file1Content = `
module MTTD
{
	struct A
	{
		0 optional B a;
	};
};
	`
	os.WriteFile("tmp/1.sdp", []byte(file1Content), os.ModePerm)
	OnFileChanged("tmp/1.sdp", file1Content)

	res, err = CheckSyntax("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range res {
		fmt.Println(v.Message)
	}

	fmt.Println("file changed2!!!")

	// 重新改正确
	file1Content = `
	#include "2.sdp"
	module MTTD
{
	struct A
	{
		0 optional B a;
	};
};
	`
	os.WriteFile("tmp/1.sdp", []byte(file1Content), os.ModePerm)
	OnFileChanged("tmp/1.sdp", file1Content)

	res, err = CheckSyntax("tmp/1.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range res {
		fmt.Println(v.Message)
	}
}

func TestCompletion() {
	content, err := os.ReadFile("test_sdp_files/completion.sdp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	OnOpenFile("test_sdp_files/completion.sdp", string(content))

	res, err := Completion("test_sdp_files/completion.sdp", 9, 19)
	fmt.Println(res, err)
}
