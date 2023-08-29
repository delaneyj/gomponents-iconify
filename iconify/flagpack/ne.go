package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackNe0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackNe2)"><use href="#flagpackNe0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackNe1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackNe1)"><path fill="#FC6500" d="M0 0v8h32V0H0Z"/><path fill="#5EAA22" d="M0 16v8h32v-8H0Z"/><path fill="#FC6500" d="M16 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></g><defs><clipPath id="flagpackNe2"><use href="#flagpackNe0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}