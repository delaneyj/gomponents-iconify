package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackFi0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackFi2)"><use href="#flagpackFi0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackFi1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackFi1)"><path fill="#2E42A5" stroke="#2E42A5" stroke-width="2" d="M10.819 1h-1v9.323H-1v4H9.819V25h4V14.323H33v-4H13.819V1h-3Z"/></g></g><defs><clipPath id="flagpackFi2"><use href="#flagpackFi0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}