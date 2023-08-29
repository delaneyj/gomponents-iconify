package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBf0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBf2)"><use href="#flagpackBf0"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackBf1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackBf1)"><path fill="#C51918" d="M0 0v12h32V0H0Z"/><path fill="#FECA00" d="m16.035 15.77l-4.703 3.56l1.505-5.797l-4.41-3.53h5.257l2.35-5.145l2.352 5.146h5.257l-4.457 3.534l1.551 5.792l-4.702-3.56Z"/></g></g><defs><clipPath id="flagpackBf2"><use href="#flagpackBf0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}