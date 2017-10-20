// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// TrElem is the React element definition corresponding to the HTML <tr> element
type TrElem struct {
	Element
}

// _TrProps defines the properties for the <tr> element
type _TrProps struct {
	*BasicHTMLElement
}

// Tr creates a new instance of a <tr> element with the provided props and
// child
func Tr(props *TrProps, children ...Element) *TrElem {

	rProps := &_TrProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &TrElem{
		Element: createElement("tr", rProps, children...),
	}
}
