package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Co(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCo0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCo2)"><use href="#flagpackCo0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackCo1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackCo1)"><path fill="#FECA00" d="M0 0v12h32V0H0Z"/><path fill="#E31D1C" d="M0 18v6h32v-6H0Z"/></g></g><defs><clipPath id="flagpackCo2"><use href="#flagpackCo0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}