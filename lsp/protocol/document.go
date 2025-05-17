package protocol

// TextDocumentIdentifier is used to identify a specific text document.
// It only contains the URI of the document.
type TextDocumentIdentifier struct {
	// URI is the unique resource identifier of the document that was closed.
	URI string `json:"uri"`
}

// VersionedTextDocumentIdentifier identifies a versioned text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// Version is the version number of the document.
	Version int `json:"version"`
}

// TextDocumentItem represents the information related to a text document.
type TextDocumentItem struct {
	// URI is the unique resource identifier of the document (usually a file path or URL).
	URI string `json:"uri"`

	// LanguageID is the language identifier associated with the document (e.g., "python", "javascript").
	LanguageID LanguageID `json:"languageId"`

	// Version is the version number of the document.
	Version int `json:"version"`

	// Text is the content of the document.
	Text string `json:"text"`
}

// TextDocumentPositionParams represents parameters for requests that operate on a specific text document
// at a specific position, such as hover information or code actions.
type TextDocumentPositionParams struct {
	// TextDocument holds the identifier of the text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// Position specifies the position within the text document.
	Position Position `json:"position"`
}

// A text edit applicable to a text document.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textEdit
type TextEdit struct {
	// The range of the text document to be manipulated. To insert text into a document create a range
	// where start == end.
	Range Range `json:"range"`

	// The string to be inserted. For delete operations use an empty string.
	NewText string `json:"newText"`
}

// TextDocumentEdit represents edits to a single text document.
type TextDocumentEdit struct {
	// TextDocument identifies the text document to change.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	// Edits is an array of edits to apply to the text document.
	Edits []TextEdit `json:"edits"`
}
