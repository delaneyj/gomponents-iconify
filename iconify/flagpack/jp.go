package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackJp0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackJp2)"><use href="#flagpackJp0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackJp1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackJp1)"><path fill="#E31D1C" fill-rule="evenodd" d="M16 19.5a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Z" clip-rule="evenodd"/></g></g><defs><clipPath id="flagpackJp2"><use href="#flagpackJp0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}