package token

const (
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// Identifiers and Literals
	IDENT = "IDENT" // user defined variable names
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Special
	ILLEGAL = "ILLEGAL" // unknown token/character
	EOF     = "EOF"     // "end of file" signals parser to stop
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
