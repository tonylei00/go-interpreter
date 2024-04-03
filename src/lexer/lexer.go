package lexer

/*
readPosition points to the “next” character in the input.
position points to the character in the input that corresponds to the ch byte
*/
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	// EOF: set ch to 0 (ASCII value for NULL)
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
