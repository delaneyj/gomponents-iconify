package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMu0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMu2)"><use href="#flagpackMu0"/><path fill="#579D20" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackMu1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackMu1)"><path fill="#FECA00" d="M0 12v6h32v-6H0Z"/><path fill="#3D58DB" d="M0 6v6h32V6H0Z"/><path fill="#E11C1B" d="M0 0v6h32V0H0Z"/></g></g><defs><clipPath id="flagpackMu2"><use href="#flagpackMu0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}