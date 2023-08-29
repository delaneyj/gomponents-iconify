package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCz0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCz2)"><use href="#flagpackCz0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackCz1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g mask="url(#flagpackCz1)"><path fill="#F7FCFF" fill-rule="evenodd" d="M0-2v14h32V-2H0Z" clip-rule="evenodd"/></g><path fill="#3D58DB" fill-rule="evenodd" d="M0 0v24l18-12L0 0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCz2"><use href="#flagpackCz0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}