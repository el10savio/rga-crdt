package rga

// Element ...
type Element struct {
	Value      string
	Identifier float64
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
func (sequence *Sequence) Decompose(element Element) (string, float64) {
	// Handle Element Not Present
	if !sequence.Lookup(element) {
		return "", -1
	}
	return element.Value, element.Identifier
}

// Before ...
func (sequence *Sequence) Before(elementPrevious, elementNext Element) bool {
	// Handle Elements Not Present
	if !sequence.Lookup(elementPrevious) || !sequence.Lookup(elementNext) {
		return false
	}

	_, identifierPrevious := sequence.Decompose(elementPrevious)
	_, identifierNext := sequence.Decompose(elementNext)
	return identifierPrevious < identifierNext
}

// AllocateIdentifierBetween ...
func AllocateIdentifierBetween(identifier1, identifier2 float64) float64 {
	// Handle identifier1 > identifier2
	if identifier1 > identifier2 {
		return -1
	}
	return (identifier1 + identifier2) / 2
}

// AddBetween ...
func (sequence *Sequence) AddBetween(value string, elementPrevious, elementNext Element) Element {
	// Handle Before(elementPrevious, elementNext) is false
	if !sequence.Before(elementPrevious, elementNext) {
		return Element{}
	}

	_, identifierPrevious := sequence.Decompose(elementPrevious)
	_, identifierNext := sequence.Decompose(elementNext)

	element := Element{Value: value, Identifier: AllocateIdentifierBetween(identifierPrevious, identifierNext)}

	// TODO: Add element to Sequence.Set
	// TODO: Downstream

	return element
}

// Remove ...
func (sequence *Sequence) Remove(element Element) {
	// Handle Element Not Present
	if !sequence.Lookup(element) {
		return
	}

	// TODO: Remove element from Sequence.Set
	// TODO: Downstream
}
