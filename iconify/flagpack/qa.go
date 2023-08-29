package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Qa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackQa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackQa1)"><use href="#flagpackQa0"/><path fill="#B61C49" d="M0 0h32v24H0z"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h6.8L12 2L6.8 4L12 6L6.8 8l5.2 2l-5.2 2l5.2 2l-5.2 2l5.2 2l-5.2 2l5.2 2l-5.2 2H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackQa1"><use href="#flagpackQa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}