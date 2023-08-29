package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTh0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTh1)"><use href="#flagpackTh0"/><path fill="#F50101" fill-rule="evenodd" d="M0 16h32v8H0v-8ZM0 0h32v6H0V0Z" clip-rule="evenodd"/><path fill="#3D58DB" stroke="#fff" stroke-width="3" d="M0 6.5h-1.5v11h35v-11H0Z"/></g><defs><clipPath id="flagpackTh1"><use href="#flagpackTh0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}