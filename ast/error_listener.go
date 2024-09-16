package ast

import (
	ant "github.com/antlr4-go/antlr/v4"
	"github.com/isolk/go-lsp/lsp/defines"
)

type SdpErrorListener struct {
	errors []defines.Diagnostic
}

var Severity = defines.DiagnosticSeverityError

func (s *SdpErrorListener) SyntaxError(recognizer ant.Recognizer, offendingSymbol interface{}, line, column int, msg string, e ant.RecognitionException) {
	chLen := 1
	token, ok := offendingSymbol.(ant.Token)
	if ok {
		chLen = len(token.GetText())
	}

	s.errors = append(s.errors, defines.Diagnostic{
		Range: defines.Range{
			Start: defines.Position{
				Line:      uint(line),
				Character: uint(column),
			},
			End: defines.Position{
				Line:      uint(line),
				Character: uint(column + chLen),
			},
		},
		Severity: &Severity,
		Source:   nil,
		Message:  msg,
	})
}

func (s *SdpErrorListener) ReportAmbiguity(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex int, exact bool, ambigAlts *ant.BitSet, configs *ant.ATNConfigSet) {
}

func (s *SdpErrorListener) ReportAttemptingFullContext(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex int, conflictingAlts *ant.BitSet, configs *ant.ATNConfigSet) {
}

func (s *SdpErrorListener) ReportContextSensitivity(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex, prediction int, configs *ant.ATNConfigSet) {
}
