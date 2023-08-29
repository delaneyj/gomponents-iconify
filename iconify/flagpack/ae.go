package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ae(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackAe0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackAe1)"><use href="#flagpackAe0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#E31D1C" d="M0 0h10v24H0z"/></g><defs><clipPath id="flagpackAe1"><use href="#flagpackAe0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}