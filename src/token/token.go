package token

const (
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// Identifiers and Literals
	IDENT = "IDENT" // variable names
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
