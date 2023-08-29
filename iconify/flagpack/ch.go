package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCh0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCh2)"><use href="#flagpackCh0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackCh1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackCh1)"><path fill="#F1F9FF" fill-rule="evenodd" d="M18 6h-4v4h-4v4h4v4h4v-4h4v-4h-4V6Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackCh2"><use href="#flagpackCh0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}