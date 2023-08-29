package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackGw0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackGw1)"><use href="#flagpackGw0"/><path fill="#FBCD17" fill-rule="evenodd" d="M16 0h16v12H16V0Z" clip-rule="evenodd"/><path fill="#0B9E7A" fill-rule="evenodd" d="M16 12h16v12H16V12Z" clip-rule="evenodd"/><path fill="#E11C1B" fill-rule="evenodd" d="M0 0h16v24H0V0Z" clip-rule="evenodd"/><path fill="#1D1D1D" fill-rule="evenodd" d="m8.93 14.606l-3.485 2.418l1.114-4.141L4 10.238l3.465-.143L8.93 6l1.466 4.095h3.458l-2.553 2.788l1.279 3.897l-3.65-2.174Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackGw1"><use href="#flagpackGw0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}