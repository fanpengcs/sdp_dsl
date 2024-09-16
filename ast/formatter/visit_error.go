package formatter

import (
	"fmt"

	ant "github.com/antlr4-go/antlr/v4"
	"github.com/isolk/go-lsp/lsp/defines"
)

type SdpErrorListener struct {
	hasError bool
	error    error
}

var Severity = defines.DiagnosticSeverityError

func (s *SdpErrorListener) SyntaxError(recognizer ant.Recognizer, offendingSymbol interface{}, line, column int, msg string, e ant.RecognitionException) {
	s.hasError = true
	s.error = fmt.Errorf("line %d:%d %s", line, column, msg)
}

func (s *SdpErrorListener) ReportAmbiguity(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex int, exact bool, ambigAlts *ant.BitSet, configs *ant.ATNConfigSet) {
}

func (s *SdpErrorListener) ReportAttemptingFullContext(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex int, conflictingAlts *ant.BitSet, configs *ant.ATNConfigSet) {
}

func (s *SdpErrorListener) ReportContextSensitivity(recognizer ant.Parser, dfa *ant.DFA, startIndex, stopIndex, prediction int, configs *ant.ATNConfigSet) {
}
