package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ng(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackNg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackNg1)"><use href="#flagpackNg0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#093" fill-rule="evenodd" d="M21 0h11v24H21V0ZM0 0h11v24H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackNg1"><use href="#flagpackNg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}