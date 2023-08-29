package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackKw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackKw2)"><use href="#flagpackKw0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackKw1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackKw1)"><path fill="#093" d="M0 0v8h32V0H0Z"/><path fill="#E31D1C" d="M0 16v8h32v-8H0Z"/></g><path fill="#272727" fill-rule="evenodd" d="M0 0v24l12-8V8L0 0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackKw2"><use href="#flagpackKw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}