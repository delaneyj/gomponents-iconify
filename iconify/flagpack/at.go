package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func At(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackAt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackAt2)"><use href="#flagpackAt0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackAt1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#C51918" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackAt1)"><path d="M0 0v8h32V0H0Zm0 16v8h32v-8H0Z"/></g></g><defs><clipPath id="flagpackAt2"><use href="#flagpackAt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}