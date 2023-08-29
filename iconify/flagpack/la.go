package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func La(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLa1)"><use href="#flagpackLa0"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 8h32v8H0V8Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M16 15.74a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLa1"><use href="#flagpackLa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}