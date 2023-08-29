package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletHalfFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m12 2l.07.003a2.41 2.41 0 0 1 1.825.907l.108.148l4.92 7.306c1.952 3.267 1.191 7.42-1.796 9.836c-2.968 2.402-7.285 2.402-10.254 0c-2.917-2.36-3.711-6.376-1.901-9.65l.134-.232l4.893-7.26c.185-.275.426-.509.709-.686a2.426 2.426 0 0 1 1.066-.36L12 2zm-1 3.149l-4.206 6.24c-1.44 2.41-.88 5.463 1.337 7.257A6.101 6.101 0 0 0 11 19.922V5.149z"/></g>`),
		g.Group(children),
	)
}