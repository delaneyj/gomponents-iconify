package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLt2)"><use href="#flagpackLt0"/><path fill="#55BA07" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackLt1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackLt1)"><path fill="#FECA00" d="M0 0v8h32V0H0Z"/><path fill="#C51918" d="M0 16v8h32v-8H0Z"/></g></g><defs><clipPath id="flagpackLt2"><use href="#flagpackLt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}