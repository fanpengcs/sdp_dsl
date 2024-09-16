package ast

import (
	"github.com/isolk/go-lsp/lsp/defines"
)

// 定义与lsp通信的api接口，实现在sdp_manager.go中

// 全局的sdp管理器
var manager = NewSdpManager()

// 格式化代码
func FormatCode(fileName string) (string, error) {
	return manager.FormatCode(fileName)
}

// 检查语法错误
func CheckSyntax(fileName string) ([]defines.Diagnostic, error) {
	res, err := manager.CheckSyntax(fileName)
	if res == nil || err != nil {
		return nil, err
	}

	tmpRes := make([]defines.Diagnostic, len(res))
	for i := 0; i < len(tmpRes); i++ {
		tmpRes[i] = res[i]
		rangeTranser(&tmpRes[i].Range)
	}
	return tmpRes, err
}

// 获取语义token
func SemanticToken(fileName string) ([]uint, error) {
	return manager.SemanticToken(fileName)
}

// 获取悬停信息
func Hover(fileName string, line, character uint) (string, error) {
	line++
	return manager.Hover(fileName, line, character)
}

// 获取定义
func Definition(fileName string, line, character uint) (*[]defines.LocationLink, error) {
	line++
	res, err := manager.Definition(fileName, line, character)
	if res == nil || err != nil {
		return nil, err
	}

	tmpRes := make([]defines.LocationLink, len(*res))
	for i := 0; i < len(*res); i++ {
		tmpRes[i] = (*res)[i]
		rangeTranser(&tmpRes[i].TargetRange)
		rangeTranser(&tmpRes[i].TargetSelectionRange)
	}
	return &tmpRes, err
}

func Completion(fileName string, line, character uint) ([]defines.CompletionItem, error) {
	line++
	res, err := manager.Completion(fileName, line, character)
	if res == nil || err != nil {
		return nil, err
	}
	return res, err
}

// 打开文件时触发
func OnOpenFile(fileName, content string) {
	manager.OnOpenFile(fileName, content)
}

// 文件内容变化时触发
func OnFileChanged(fileName, content string) {
	manager.OnFileChanged(fileName, content)
}

// 删除文件时触发
func OnDeleteFile(fileName string) {
	manager.OnDeleteFile(fileName)
}

// 重命名文件时触发
func OnRenameFile(oldName, newName string) {
	manager.OnRenameFile(oldName, newName)
}

func rangeTranser(s *defines.Range) {
	s.Start.Line--
	s.End.Line--
}
