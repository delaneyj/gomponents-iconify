package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGm1)"><use href="#flagpackGm0"/><path fill="#5EAA22" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#E31D1C" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/><path fill="#3D58DB" stroke="#fff" stroke-width="3" d="M0 8.5h-1.5v7h35v-7H0Z"/></g><defs><clipPath id="flagpackGm1"><use href="#flagpackGm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}