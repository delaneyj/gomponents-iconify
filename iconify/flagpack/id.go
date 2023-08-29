package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Id(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackId0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackId2)"><use href="#flagpackId0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackId1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackId1)"><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v12h32V0H0Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackId2"><use href="#flagpackId0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}