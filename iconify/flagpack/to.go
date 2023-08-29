package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func To(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTo0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTo2)"><use href="#flagpackTo0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTo1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackTo1)"><path fill="#F7FCFF" d="M0 0h18v16H0z"/><path fill="#E31D1C" fill-rule="evenodd" d="M12 2H8v4H4v4h4v4h4v-4h4V6h-4V2Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackTo2"><use href="#flagpackTo0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}