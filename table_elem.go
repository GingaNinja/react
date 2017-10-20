// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// TableElem is the React element definition corresponding to the HTML <table> element
type TableElem struct {
	Element
}

// _TableProps defines the properties for the <table> element
type _TableProps struct {
	*BasicHTMLElement
}

// Table creates a new instance of a <table> element with the provided props and
// child
func Table(props *TableProps, children ...Element) *TableElem {

	rProps := &_TableProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &TableElem{
		Element: createElement("table", rProps, children...),
	}
}
