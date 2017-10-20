// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// ThElem is the React element definition corresponding to the HTML <th> element
type ThElem struct {
	Element
}

// _ThProps defines the properties for the <th> element
type _ThProps struct {
	*BasicHTMLElement
}

// Th creates a new instance of a <th> element with the provided props and
// child
func Th(props *ThProps, children ...Element) *ThElem {

	rProps := &_ThProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &ThElem{
		Element: createElement("th", rProps, children...),
	}
}
