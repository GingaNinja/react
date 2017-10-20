// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// TheadElem is the React element definition corresponding to the HTML <thead> element
type TheadElem struct {
	Element
}

// _TheadProps defines the properties for the <thead> element
type _TheadProps struct {
	*BasicHTMLElement
}

// Thead creates a new instance of a <thead> element with the provided props and
// child
func Thead(props *TheadProps, children ...Element) *TheadElem {

	rProps := &_TheadProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &TheadElem{
		Element: createElement("thead", rProps, children...),
	}
}
