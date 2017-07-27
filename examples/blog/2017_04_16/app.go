// Template generated by reactGen

package main

import (
	r "myitcv.io/react"
)

type AppDef struct {
	r.ComponentDef
}

func App() *AppElem {
	return buildAppElem()
}

func (a AppDef) Render() r.Element {
	return r.Div(nil,
		r.H1(nil,
			r.S("Hello World"),
		),
		r.P(nil,
			r.S("This is my first GopherJS React App."),
		),
		FooBar(FooBarProps{Name: "Peter"}),
	)
}
