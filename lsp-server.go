package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sdp-lsp/ast"

	"github.com/isolk/go-lsp/jsonrpc"
	"github.com/isolk/go-lsp/logs"
	"github.com/isolk/go-lsp/lsp"
	"github.com/isolk/go-lsp/lsp/defines"
)

func startLspServer() {
	f, e := os.Create("/tmp/log.txt")
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	logs.Init(log.New(f, "", log.LstdFlags))

	server := lsp.NewServer(&lsp.Options{})

	server.OnInitialize(func(ctx context.Context, req *defines.InitializeParams) (result *defines.InitializeResult, err *defines.InitializeError) {
		ser := defines.ServerCapabilities{}
		ser.TextDocumentSync = defines.TextDocumentSyncKindFull
		ser.PositionEncoding = defines.PositionEncodingKind_UTF8

		ser.HoverProvider = true              // hover
		ser.DefinitionProvider = true         // go to definition
		ser.DocumentFormattingProvider = true // format code
		ser.CompletionProvider = &defines.CompletionOptions{}
		// ser.DiagnosticProvider = true         // diagnostic
		ser.SemanticTokensProvider = &defines.SemanticTokensOptions{ // semantic token
			Legend: defines.SemanticTokensLegend{
				TokenTypes:     []string{"namespace", "type", "struct", "typeParameter", "parameter", "variable", "property", "enumMember", "event", "function", "method", "macro", "keyword", "modifier", "comment", "string", "number", "regexp", "operator", "decorator"},
				TokenModifiers: []string{"declaration", "definition", "readonly", "static", "deprecated", "abstract", "async", "modification", "documentation", "defaultLibrary"},
			},
			Full: &Ture,
		}
		result = &defines.InitializeResult{
			Capabilities: ser,
		}
		return
	})

	server.OnHover(func(ctx context.Context, req *defines.HoverParams) (result *defines.Hover, err error) {
		fileName := uriToFileName(req.TextDocument.Uri)
		line := req.Position.Line
		colume := req.Position.Character
		res, err := ast.Hover(fileName, line, colume)
		if err != nil || res == "" {
			return nil, err
		}

		hover := &defines.Hover{
			Contents: defines.MarkupContent{
				Kind:  defines.MarkupKindMarkdown,
				Value: "```c++\n" + res + "\n```",
			},
		}
		return hover, nil
	})

	server.OnDefinition(func(ctx context.Context, req *defines.DefinitionParams) (result *[]defines.LocationLink, err error) {
		fileName := uriToFileName(req.TextDocument.Uri)
		line := req.Position.Line
		colume := req.Position.Character
		res, err := ast.Definition(fileName, line, colume)
		if err != nil {
			return nil, err
		}
		return res, nil
	})

	server.OnDocumentFormatting(func(ctx context.Context, req *defines.DocumentFormattingParams) (result *[]defines.TextEdit, err error) {
		res, err := ast.FormatCode(uriToFileName(req.TextDocument.Uri))
		if err != nil {
			return nil, err
		}
		if res == "" {
			return nil, nil
		}

		result = &[]defines.TextEdit{
			{
				Range: defines.Range{
					Start: defines.Position{Line: 0, Character: 0},
					End:   defines.Position{Line: 10000000, Character: 10000000},
				},
				NewText: res,
			},
		}
		return result, nil
	})

	// server.OnDocumentDiagnostic(func(ctx context.Context, req *defines.DocumentDiagnosticParams) (result defines.DocumentDiagnosticReport, err error) {
	// 	fileName := uriToFileName(req.TextDocument.Uri)
	// 	items, err := ast.CheckSyntax(fileName)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	kind := defines.DocumentDiagnosticReportKindFull
	// 	res := defines.RelatedFullDocumentDiagnosticReport{
	// 		FullDocumentDiagnosticReport: defines.FullDocumentDiagnosticReport{
	// 			Kind:  &kind,
	// 			Items: items,
	// 		},
	// 	}
	// 	return res, nil
	// })

	server.OnDocumentSemanticTokenFull(func(ctx context.Context, req *defines.SemanticTokensParams) (result defines.SemanticTokens, err error) {
		fileName := uriToFileName(req.TextDocument.Uri)
		items, err := ast.SemanticToken(fileName)
		if err != nil {
			return result, err
		}
		result = defines.SemanticTokens{
			Data: items,
		}
		return
	})

	server.OnCompletion(func(ctx context.Context, req *defines.CompletionParams) (result *[]defines.CompletionItem, err error) {
		fileName := uriToFileName(req.TextDocument.Uri)
		line := req.Position.Line
		colume := req.Position.Character
		// content := *req.Context.TriggerCharacter
		res, err := ast.Completion(fileName, line, colume)
		if err != nil {
			return nil, err
		}
		return &res, nil
	})

	server.OnDidOpenTextDocument(func(ctx context.Context, req *defines.DidOpenTextDocumentParams) (err error) {
		ast.OnOpenFile(uriToFileName(req.TextDocument.Uri), req.TextDocument.Text)

		session := jsonrpc.GetSession(ctx)
		if session == nil {
			logs.Println("session is nil")
			return
		}
		fileName := uriToFileName(req.TextDocument.Uri)
		items, err := ast.CheckSyntax(fileName)
		if err != nil {
			logs.Println(err)
			return
		}
		if len(items) == 0 {
			items = make([]defines.Diagnostic, 0)
		}
		session.SendNotification(defines.MethodName_PublishDiagnostics, defines.PublishDiagnosticsParams{Uri: req.TextDocument.Uri, Diagnostics: items})
		return
	})

	server.OnDidChangeTextDocument(func(ctx context.Context, req *defines.DidChangeTextDocumentParams) (err error) {
		fileName := uriToFileName(req.TextDocument.Uri)
		for _, v := range req.ContentChanges {
			ast.OnFileChanged(fileName, v.Text.(string))
		}
		return
	})

	server.OnDidSaveTextDocument(func(ctx context.Context, req *defines.DidSaveTextDocumentParams) (err error) {
		session := jsonrpc.GetSession(ctx)
		if session == nil {
			logs.Println("session is nil")
			return
		}
		fileName := uriToFileName(req.TextDocument.Uri)
		items, err := ast.CheckSyntax(fileName)
		if err != nil {
			logs.Println(err)
			return
		}
		if len(items) == 0 {
			items = make([]defines.Diagnostic, 0)
		}
		session.SendNotification(defines.MethodName_PublishDiagnostics, defines.PublishDiagnosticsParams{Uri: req.TextDocument.Uri, Diagnostics: items})
		return
	})

	server.Run()
}

func uriToFileName(uri defines.DocumentUri) string {
	if len(uri) < 7 {
		return ""
	}

	return string(uri[7:])
}
