package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackMm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackMm1)"><use href="#flagpackMm0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 8h32v8H0V8Z" clip-rule="evenodd"/><path fill="#FFD018" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m16.062 15.976l-5.15 3.274l1.727-5.733l-3.674-3.745l5.065-.11l2.241-5.66l2.042 5.734l5.053.088l-3.797 3.814l1.773 5.454l-5.28-3.116Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackMm1"><use href="#flagpackMm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}