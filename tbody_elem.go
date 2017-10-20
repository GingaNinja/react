// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package react

// TbodyElem is the React element definition corresponding to the HTML <tbody> element
type TbodyElem struct {
	Element
}

// _TbodyProps defines the properties for the <tbody> element
type _TbodyProps struct {
	*BasicHTMLElement
}

// Tbody creates a new instance of a <tbody> element with the provided props and
// child
func Tbody(props *TbodyProps, children ...Element) *TbodyElem {

	rProps := &_TbodyProps{
		BasicHTMLElement: newBasicHTMLElement(),
	}

	if props != nil {
		props.assign(rProps)
	}

	return &TbodyElem{
		Element: createElement("tbody", rProps, children...),
	}
}
