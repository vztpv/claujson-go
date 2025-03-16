package claujson

type Token struct {
	// union of int64, float64, bool, string(key or not), null, start_array, start_object, end_array, end_object
}

// handler
type Scanner struct {
	tok []Token
}

func (h *Scanner) Null() {
	//h.buf = append(h.buf, "null "...)
}

func (h *Scanner) Bool(v bool) {
	//h.buf = append(h.buf, fmt.Sprintf("%t ", v)...)
}

func (h *Scanner) Int(v int64) {
	//h.buf = append(h.buf, fmt.Sprintf("%d ", v)...)
}

func (h *Scanner) Float(v float64) {
	//h.buf = append(h.buf, fmt.Sprintf("%g ", v)...)
}

func (h *Scanner) Number(v string) {
	//h.buf = append(h.buf, fmt.Sprintf("%s ", v)...)
}

func (h *Scanner) String(v string) {
	//h.buf = append(h.buf, fmt.Sprintf("%s ", v)...)
}

func (h *Scanner) ObjectStart() {
	//h.buf = append(h.buf, '{')
	//h.buf = append(h.buf, ' ')
}

func (h *Scanner) ObjectEnd() {
	//h.buf = append(h.buf, '}')
	//h.buf = append(h.buf, ' ')
}

func (h *Scanner) Key(v string) {
	//h.buf = append(h.buf, fmt.Sprintf("%s: ", v)...)
}

func (h *Scanner) ArrayStart() {
	//h.buf = append(h.buf, '[')
	//h.buf = append(h.buf, ' ')
}

func (h *Scanner) ArrayEnd() {
	//h.buf = append(h.buf, ']')
	//h.buf = append(h.buf, ' ')
}
