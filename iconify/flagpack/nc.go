package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackNc0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackNc1)"><use href="#flagpackNc0"/><path fill="#F50100" fill-rule="evenodd" d="M22 0h10v24H22V0Z" clip-rule="evenodd"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0h12v24H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M10 0h12v24H10V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackNc1"><use href="#flagpackNc0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}