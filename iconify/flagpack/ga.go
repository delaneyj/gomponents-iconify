package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ga(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGa0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGa1)"><use href="#flagpackGa0"/><path fill="#40A8FF" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#FECA00" fill-rule="evenodd" d="M0 8h32v8H0V8Z" clip-rule="evenodd"/><path fill="#73BE4A" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGa1"><use href="#flagpackGa0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}