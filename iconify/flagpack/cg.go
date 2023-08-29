package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackCg0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackCg1)"><use href="#flagpackCg0"/><path fill="#FA1111" fill-rule="evenodd" d="M32 0v24H0L32 0Z" clip-rule="evenodd"/><path fill="#07A907" fill-rule="evenodd" d="M0 24V0h32L0 24Z" clip-rule="evenodd"/><path fill="#FBCD17" fill-rule="evenodd" d="M29.492-5.8L-1 23.576l6.052 3.012L34.64-.212L29.49-5.8Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackCg1"><use href="#flagpackCg0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}