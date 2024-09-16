package ast

import (
	"errors"
	"strings"

	"github.com/isolk/go-lsp/lsp/defines"
)

type SdpFileManager struct {
	files map[string]*SdpFile
}

func NewSdpFileManager() *SdpFileManager {
	return &SdpFileManager{
		files: map[string]*SdpFile{},
	}
}

// 添加新的sdp文件，并且会根据include情况，引进新的文件
func (s *SdpFileManager) add(fileName, content string) (file *SdpFile, err error) {
	sdpFile := s.files[fileName]
	if sdpFile != nil {
		return nil, errors.New("file already exists")
	}

	file = NewSdpFile(fileName, content)
	s.files[fileName] = file
	file.parse()
	return
}

func (s *SdpFileManager) update(fileName, content string) (file *SdpFile, err error) {
	sdpFile := s.files[fileName]
	if sdpFile == nil {
		return nil, errors.New("file not exists")
	}

	file = NewSdpFile(fileName, content)
	s.files[fileName] = file
	file.parse()
	return
}

func (s *SdpFileManager) delete(fileName string) error {
	delete(s.files, fileName)
	return nil
}

func (s *SdpFileManager) has(fileName string) (has bool) {
	return
}

func (s *SdpFileManager) get(fileName string) (file *SdpFile) {
	file = s.files[fileName]
	return
}

func (s *SdpFileManager) reparseError(fileList []string) {
	for _, v := range fileList {
		file := s.files[v]
		if file == nil {
			continue
		}
		file.parse()
	}
}

func (s *SdpFileManager) searchType(module, typeName string) string {
	for _, v := range s.files {
		res := v.searchType(module, typeName)
		if res != "" {
			return res
		}
	}
	return ""
}

func (s *SdpFileManager) searchTypeLocation(module, typeName string) []defines.LocationLink {
	for _, v := range s.files {
		if res := v.searchTypeLocation(module, typeName); len(res) > 0 {
			return res
		}
	}
	return nil
}

type TypeContent struct {
	Content  string
	Module   string
	FileName string
}

func (s *SdpFileManager) searchTypeByContent(content string) []TypeContent {
	res := []TypeContent{}
	for _, v := range s.files {
		allTypes := v.getAllTypes()
		allTypes = find(allTypes, content)
		res = append(res, allTypes...)
	}
	return res
}

func find(targetList []TypeContent, target string) []TypeContent {
	tmp := []TypeContent{}
	for _, v := range targetList {
		// tmp[v] = strings.ToLower(v)
		// tmp[v] = strings.ToLower(v)
		v.Content = strings.ToLower(v.Content)
		tmp = append(tmp, v)
	}
	target = strings.ToLower(target)

	// 模糊查询，只要targetList里面包含target的全部字符即可，不要求顺序
	res := []TypeContent{}
	for i, t := range tmp {
		if fuzzyContains(t.Content, target) {
			res = append(res, targetList[i])
		}
	}

	return res
}

func findString(targetList []string, target string) []string {
	tmp := []string{}
	for _, v := range targetList {
		tmp = append(tmp, strings.ToLower(v))
	}
	target = strings.ToLower(target)

	// 模糊查询，只要targetList里面包含target的全部字符即可，不要求顺序
	res := []string{}
	for i, t := range tmp {
		if fuzzyContains(t, target) {
			res = append(res, targetList[i])
		}
	}
	return res
}

func fuzzyContains(a, b string) bool {
	for _, char := range b {
		if strings.Count(a, string(char)) < strings.Count(b, string(char)) {
			return false
		}
	}
	return true
}
