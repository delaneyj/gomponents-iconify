package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func It(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackIt0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackIt1)"><use href="#flagpackIt0"/><path fill="#C51918" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0h12v24H0V0Z" clip-rule="evenodd"/><path fill="#fff" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackIt1"><use href="#flagpackIt0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}