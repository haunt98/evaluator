package scanner

// Wrap Scanner with a buffer
type BufferScanner struct {
	s   *Scanner
	buf buffer
}

type buffer struct {
	tokenText TokenText
	isCached  bool
}

func NewBufferScanner(s *Scanner) *BufferScanner {
	return &BufferScanner{
		s: s,
	}
}

// Scan() return next token and it's gone
func (bs *BufferScanner) Scan() TokenText {
	if bs.buf.isCached {
		bs.buf.isCached = false
		return bs.buf.tokenText
	}

	bs.buf.tokenText = bs.s.Scan()
	return bs.buf.tokenText
}

// Peek() return next token but it's still there
func (bs *BufferScanner) Peek() TokenText {
	tokenText := bs.Scan()
	bs.undoScan()
	return tokenText
}

func (bs *BufferScanner) undoScan() {
	bs.buf.isCached = true
}
