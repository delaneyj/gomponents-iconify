package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGn0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGn1)"><use href="#flagpackGn0"/><path fill="#FBCD17" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/><path fill="#0B9E7A" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M0 0h10v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGn1"><use href="#flagpackGn0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}