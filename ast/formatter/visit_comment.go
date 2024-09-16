package formatter

import (
	"fmt"
	"sdp-lsp/ast/parsing"
	"strings"

	ant "github.com/antlr4-go/antlr/v4"
)

// 如果有注释，就返回注释，否则返回空字符串
func (v *SdpVisitor) VisitBeforComment(token ant.Token, stream ant.TokenStream, spaceNum int) string {
	if v.hasError {
		panic("has error")
	}

	if token.GetTokenIndex() <= 0 {
		return ""
	}

	// 从当前token往前找到第一个非隐藏的字符，这时候就是注释的开始，或者是换行符
	startIndex := token.GetTokenIndex() - 1
	for ; startIndex >= 0; startIndex-- {
		chType := stream.Get(startIndex).GetTokenType()
		if chType != parsing.SdpLINE_COMMENT && chType != parsing.SdpNEW_LINE && chType != parsing.SdpBLOCK_COMMENT {
			startIndex++
			break
		}
	}

	if startIndex < 0 {
		startIndex = 0
	}

	res := ""
	lineCount := 0
	for i := startIndex; i < token.GetTokenIndex(); i++ {
		ch := stream.Get(i)
		chType := ch.GetTokenType()
		if chType == parsing.SdpNEW_LINE {
			lineCount++
			continue
		}

		switch {
		case lineCount == 0 && i != 0:
			continue // 表示是单行注释，由上一个处理
		case lineCount == 1:
			res += "\n"
		case lineCount > 1:
			res += "\n\n"
		}

		if chType == parsing.SdpBLOCK_COMMENT {
			tx := ch.GetText()
			strs := strings.Split(tx, "\n")
			for i, subStr := range strs {
				res += replaceLeftSpace(subStr, spaceNum)
				if i != len(strs)-1 {
					res += "\n"
				}
			}

		}

		if chType == parsing.SdpLINE_COMMENT {
			res += fmt.Sprintf("%v%v", strings.Repeat(" ", spaceNum), ch.GetText())
		}
		lineCount = 0
	}

	switch {
	case lineCount == 0:
		res += "\n"
	case lineCount == 1:
		res += "\n"
	case lineCount > 1:
		res += "\n\n"
	}

	return res
}

func (v *SdpVisitor) VisitRightComment(token ant.Token, stream ant.TokenStream) string {
	if v.hasError {
		panic("has error")
	}

	next := token.GetTokenIndex() + 1
	if next >= stream.Size() {
		return ""
	}

	ch := stream.Get(next)
	chType := ch.GetTokenType()
	if chType == parsing.SdpLINE_COMMENT || chType == parsing.SdpBLOCK_COMMENT {
		return ch.GetText() + ""
	}
	return ""
}

func replaceLeftSpace(str string, spaceNum int) string {
	// 先移除spaceNum个空格，然后判断是否需要补充空格
	trimNum := 0
	res := ""
	for _, char := range str {
		if char == ' ' || char == '\t' {
			trimNum++
			continue
		}
		res = str[trimNum:]
		break
	}

	// 移除的刚好相等
	if trimNum == spaceNum {
		return str
	} else if trimNum > spaceNum {
		return strings.Repeat(" ", trimNum-spaceNum) + res
	} else {
		return strings.Repeat(" ", spaceNum) + res
	}
}
