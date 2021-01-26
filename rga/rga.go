package rga

// Element ...
type Element struct {
	Value      string
	Identifier string
}

// Sequence ...
type Sequence struct {
	Set []Element
}

// Initialize ...
func (sequence *Sequence) Initialize() Sequence {
	return Sequence{}
}

// Lookup ...
func (sequence *Sequence) Lookup(element Element) bool {
	for _, _element := range sequence.Set {
		if _element.Identifier == element.Identifier {
			return true
		}
	}
	return false
}

// Decompose ...
func (sequence *Sequence) Decompose(element Element) (string, string) {
	// TODO: Err Handle Element Not Present
	return element.Value, element.Identifier
}

// Before ...
func (sequence *Sequence) Before(element, elementNext Element) bool {
	// TODO: Err Handle Elements Not Present
	_, identifier := sequence.Decompose(element)
	_, identifierNext := sequence.Decompose(elementNext)
	return identifier < identifierNext
}
