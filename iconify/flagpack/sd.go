package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackSd0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackSd1)"><use href="#flagpackSd0"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 0h32v24H0V0Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0v8h32V0H0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="M0 16v8h32v-8H0Z" clip-rule="evenodd"/><path fill="#5EAA22" fill-rule="evenodd" d="m0 0l16 12L0 24V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackSd1"><use href="#flagpackSd0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}