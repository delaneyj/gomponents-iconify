package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Td(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTd0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTd1)"><use href="#flagpackTd0"/><path fill="#E11C1B" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0h12v24H0V0Z" clip-rule="evenodd"/><path fill="#FBCD17" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackTd1"><use href="#flagpackTd0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}