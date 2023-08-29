package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Be(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackBe0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackBe1)"><use href="#flagpackBe0"/><path fill="#FECA00" fill-rule="evenodd" d="M10 0h11v24H10V0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M21 0h11v24H21V0Z" clip-rule="evenodd"/><path fill="#1D1D1D" fill-rule="evenodd" d="M0 0h11v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackBe1"><use href="#flagpackBe0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}