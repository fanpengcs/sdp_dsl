package ast

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/isolk/go-lsp/logs"
	"github.com/isolk/go-lsp/lsp/defines"
)

type SdpManager struct {
	includeManager *IncludeManager
	sdpFileManager *SdpFileManager

	ch chan bool
}

func NewSdpManager() *SdpManager {
	return &SdpManager{
		includeManager: NewIncludeManager(),
		sdpFileManager: NewSdpFileManager(),
		ch:             make(chan bool, 1),
	}
}

func (s *SdpManager) OnOpenFile(fileName, content string) error {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	if s.sdpFileManager.has(fileName) {
		return nil
	}

	// add在处理时，可能会根据include情况，递归add其他的文件
	sdpFile, err := s.sdpFileManager.add(fileName, content)
	if err != nil {
		return err
	}

	s.includeManager.resetRef(fileName, sdpFile.getIncludeList())
	s.includeManager.recaculateRealRef()
	return err
}

func (s *SdpManager) OnFileChanged(fileName, content string) error {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile, err := s.sdpFileManager.update(fileName, content)
	if err != nil {
		return err
	}

	compareSlices := func(slice1, slice2 []string) bool {
		if len(slice1) != len(slice2) {
			return false
		}
		for i, v := range slice1 {
			if v != slice2[i] {
				return false
			}
		}
		return true
	}

	if !compareSlices(sdpFile.getIncludeList(), s.includeManager.getRef(fileName)) {
		s.includeManager.resetRef(fileName, sdpFile.getIncludeList())
		s.includeManager.recaculateRealRef()
	}

	list := s.includeManager.getWhoRefMe(fileName)
	s.sdpFileManager.reparseError(list)

	return nil
}

func (s *SdpManager) OnDeleteFile(fileName string) error {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	s.sdpFileManager.delete(fileName)
	s.includeManager.resetRef(fileName, nil)

	list := s.includeManager.getWhoRefMe(fileName)
	s.sdpFileManager.reparseError(list)
	return nil
}

func (s *SdpManager) OnRenameFile(oldName, newName string) error {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	s.sdpFileManager.delete(oldName)
	s.includeManager.resetRef(oldName, nil)

	file, err := os.ReadFile(newName)
	if err != nil {
		return err
	}

	newSdpFile, err := s.sdpFileManager.add(newName, string(file))
	if err != nil {
		return err
	}
	s.includeManager.resetRef(newName, newSdpFile.getIncludeList())

	list := s.includeManager.getWhoRefMe(oldName)
	s.sdpFileManager.reparseError(list)

	list = s.includeManager.getWhoRefMe(oldName)
	s.sdpFileManager.reparseError(list)
	return nil
}

// 格式化代码
func (s *SdpManager) FormatCode(fileName string) (res string, err error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile := s.sdpFileManager.get(fileName)
	if sdpFile == nil {
		return "", errors.New("file not found")
	}
	return sdpFile.getFormatCode()
}

// 检查语法错误
func (s *SdpManager) CheckSyntax(fileName string) (syntaxErrors []defines.Diagnostic, err error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile := s.sdpFileManager.get(fileName)
	if sdpFile == nil {
		return nil, errors.New("file not found")
	}
	return sdpFile.getSyntaxErrors(), nil
}

// 获取语义
func (s *SdpManager) SemanticToken(fileName string) ([]uint, error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile := s.sdpFileManager.get(fileName)
	if sdpFile == nil {
		return nil, errors.New("file not found")
	}

	tokenData := sdpFile.getTokenDatas()
	sort.Slice(tokenData, func(i, j int) bool {
		if tokenData[i].Line < tokenData[j].Line {
			return true
		}
		if tokenData[i].Line == tokenData[j].Line && tokenData[i].Column < tokenData[j].Column {
			return true
		}
		return false
	})

	data := []uint{}
	for i := 0; i < len(tokenData); i++ {
		d := make([]uint, 5)
		if i == 0 {
			d[0] = uint(tokenData[i].Line - 1)
			d[1] = uint(tokenData[i].Column)
		} else if tokenData[i].Line == tokenData[i-1].Line {
			d[0] = 0
			d[1] = uint(tokenData[i].Column - tokenData[i-1].Column)
		} else {
			d[0] = uint(tokenData[i].Line - tokenData[i-1].Line)
			d[1] = uint(tokenData[i].Column)
		}
		d[2] = uint(tokenData[i].Length)
		d[3] = uint(tokenData[i].TokenType)
		d[4] = uint(tokenData[i].Modifier)
		data = append(data, d...)
	}
	return data, nil
}

func (s *SdpManager) Hover(fileName string, line, character uint) (string, error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile := s.sdpFileManager.get(fileName)
	if sdpFile == nil {
		return "", errors.New("file not found")
	}

	typeName, module := sdpFile.getHoverData(line, character)
	if typeName == "" || module == "" {
		return "", nil
	}
	content := s.sdpFileManager.searchType(module, typeName)
	return content, nil
}

func (s *SdpManager) Definition(fileName string, line, character uint) (*[]defines.LocationLink, error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))

			os.Exit(1)
		}
	}()

	SdpFile := s.sdpFileManager.get(fileName)
	if SdpFile == nil {
		return nil, errors.New("file not found")
	}

	// 获取悬停信息，从而确定当前位置的类型和模块
	typeName, module := SdpFile.getHoverData(line, character)
	if typeName != "" && module != "" {
		// 获取对应的类型定义位置
		location := s.sdpFileManager.searchTypeLocation(module, typeName)
		if location == nil {
			return nil, nil
		}
		return &location, nil
	}

	// 上面type没找到，表示当前位置不是类型。 那么再找下是不是include文件，目前只允许这两种模式的跳转
	includeFileName := SdpFile.getIncludeFileName(line, character)
	if includeFileName == "" {
		return nil, nil
	}

	location := []defines.LocationLink{{
		TargetUri: defines.DocumentUri(includeFileName),
		TargetRange: defines.Range{
			Start: defines.Position{
				Line:      1,
				Character: 0,
			},
			End: defines.Position{
				Line:      1,
				Character: 1,
			},
		},
		TargetSelectionRange: defines.Range{
			Start: defines.Position{
				Line:      1,
				Character: 0,
			},
			End: defines.Position{
				Line:      1,
				Character: 1,
			},
		},
	}}

	return &location, nil
}

func (s *SdpManager) Completion(fileName string, line, character uint) ([]defines.CompletionItem, error) {
	s.lock()
	defer s.releaseLock()
	defer func() {
		if r := recover(); r != nil {
			logs.Printf("panic!,err=%v,stack=%s:", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	sdpFile := s.sdpFileManager.get(fileName)
	if sdpFile == nil {
		return nil, errors.New("file not found")
	}

	contet, moduleName := sdpFile.getContent(line, character)
	if contet == "" {
		return nil, nil
	}

	typeList := s.sdpFileManager.searchTypeByContent(contet)
	res := make([]defines.CompletionItem, 0)
	for _, typeInfo := range typeList {
		typeFileName := findRelativeFileName(fileName, typeInfo.FileName)
		content := typeInfo.Content
		if moduleName != typeInfo.Module {
			content = typeInfo.Module + "::" + content
		}

		var textEdit *[]defines.TextEdit = nil

		if !s.getRealRef(fileName)[typeInfo.FileName] {
			textEdit = &[]defines.TextEdit{{
				Range: defines.Range{
					Start: defines.Position{
						Line:      0,
						Character: 0,
					},
					End: defines.Position{
						Line:      0,
						Character: 0,
					},
				},
				NewText: fmt.Sprintf("#include \"%s\"\n", typeFileName),
			}}
		}
		res = append(res, defines.CompletionItem{
			Label:               content,
			Detail:              &typeFileName,
			AdditionalTextEdits: textEdit,
		})
	}
	keyList := findString(keywords, contet)
	keyWord := defines.CompletionItemKindKeyword
	for _, v := range keyList {
		res = append(res, defines.CompletionItem{
			Label: v,
			Kind:  &keyWord,
		})
	}
	return res, nil
}

func findRelativeFileName(targetFileName, newFileName string) string {
	// 使用 filepath.Dir() 获取文件路径的父级目录
	dirA := filepath.Dir(targetFileName)
	dirB := filepath.Dir(newFileName)

	// 使用 filepath.Rel() 获取相对路径
	relativePath, _ := filepath.Rel(dirA, dirB)

	// 处理相对路径中的 "." 和 ".."
	// relativePath = filepath.Clean(relativePath)
	res := relativePath + "/" + filepath.Base(newFileName)
	res = strings.TrimPrefix(res, "./")
	return res
}

// 内部方法不加锁
func (s *SdpManager) getFile(fileName string) *SdpFile {
	return s.sdpFileManager.get(fileName)
}

// 内部方法不加锁
func (s *SdpManager) addFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, err = s.sdpFileManager.add(fileName, string(file))
	return err
}

// 内部方法不加锁
func (s *SdpManager) addRef(fileName, targetFileName string) error {
	s.includeManager.addRef(fileName, targetFileName)
	return nil
}

func (s *SdpManager) getRealRef(fileName string) map[string]bool {
	return s.includeManager.getRealRef(fileName)
}

func (s *SdpManager) lock() {
	s.ch <- true
}

func (s *SdpManager) releaseLock() {
	<-s.ch
}
