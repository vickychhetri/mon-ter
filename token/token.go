package token

type TokenType string

// Token represents a single token in the source code.
// It consists of a type and a literal value.
// The type indicates the kind of token (e.g., identifier, keyword, operator),
// while the literal value is the actual string representation of the token in the source code.
type Token struct {
	Type    TokenType
	Literal string
}

// Define the various token types as constants.
// These constants represent the different kinds of tokens that can be recognized by the lexer.
// Each token type is associated with a specific string value.
// For example, the token type "ILLEGAL" represents an illegal token, while "EOF" represents the end of the file.
// Other token types include identifiers, literals, operators, delimiters, and keywords.
// The keywords are mapped to their corresponding token types using a map for easy lookup.
// The keywords include "fn" for function and "let" for variable declaration.
// The LookupIdent function checks if a given identifier is a keyword and returns the corresponding token type.
// If the identifier is not a keyword, it returns the token type "IDENT".
// This allows the lexer to differentiate between keywords and regular identifiers in the source code.
// The lexer uses these token types to categorize the tokens it encounters while scanning the source code.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// keywords maps language keywords to their corresponding token types.
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookmeUp checks if the identifier is a keyword and returns the corresponding token type
// or returns IDENT if it is not a keyword.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// LookupIdent checks if the identifier is a keyword and returns the corresponding token type
// or returns IDENT if it is not a keyword.
// func LookupIdent(ident string) TokenType {
// 	switch ident {
// 	case "fn":
// 		return FUNCTION
// 	case "let":
// 		return LET
// 	default:
// 		return IDENT
// 	}

// }
