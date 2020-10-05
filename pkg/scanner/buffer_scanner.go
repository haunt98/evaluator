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
// if cached is true -> return buffer and set cached is false
// if cached is false -> return scanner result and stores it in buffer
// always stores scanner result in buffer
// but buffer is returned only 1 time, next time buffer is updated from scanner result
func (bs *BufferScanner) Scan() TokenText {
	if bs.buf.isCached {
		bs.buf.isCached = false
		return bs.buf.tokenText
	}

	bs.buf.tokenText = bs.s.Scan()
	return bs.buf.tokenText
}

// Peek() return next token but it's still there
// return scan result and set cached is true
// next time scan will get result from buffer
func (bs *BufferScanner) Peek() TokenText {
	tokenText := bs.Scan()
	bs.buf.isCached = true
	return tokenText
}
