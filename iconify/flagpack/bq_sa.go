package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BqSa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBqSa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBqSa1)"><use href="#flagpackBqSa0"/><path fill="#fff" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#F00000" fill-rule="evenodd" d="M0 11.664V0h16L0 11.664Zm32 0V0H16l16 11.664Z" clip-rule="evenodd"/><path fill="#00268D" fill-rule="evenodd" d="M0 11.664V24h16L0 11.664Zm32 0v12.672L16 24l16-12.336Z" clip-rule="evenodd"/><path fill="#FEDA00" fill-rule="evenodd" d="m15.714 14.256l-3.518 2.507l1.291-4.141L10 10.087h4.345L15.714 6l1.452 4.087h4.226l-3.457 2.535l1.314 4.141l-3.535-2.507Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBqSa1"><use href="#flagpackBqSa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}