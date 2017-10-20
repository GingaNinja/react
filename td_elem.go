// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// TdElem is the React element definition corresponding to the HTML <td> element
type TdElem struct {
	Element
}

// _TdProps defines the properties for the <td> element
type _TdProps struct {
	*BasicHTMLElement
}

// Td creates a new instance of a <td> element with the provided props and
// child
func Td(props *TdProps, children ...Element) *TdElem {

	rProps := &_TdProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &TdElem{
		Element: createElement("td", rProps, children...),
	}
}
