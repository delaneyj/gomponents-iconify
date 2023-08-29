package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16.401 20.5L6 2m16 17a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/><path fill="currentColor" d="M5 21.25a.75.75 0 0 1 0 1.5v-1.5ZM8.75 19v.75h-1.5V19h1.5Zm-.498 1.868a.75.75 0 0 1-1.307-.736l1.307.736Zm9.094-19.236a.75.75 0 0 1 1.308.736l-1.308-.736ZM2.75 19A2.25 2.25 0 0 0 5 21.25v1.5A3.75 3.75 0 0 1 1.25 19h1.5Zm4.5 0A2.25 2.25 0 0 0 5 16.75v-1.5A3.75 3.75 0 0 1 8.75 19h-1.5ZM5 16.75A2.25 2.25 0 0 0 2.75 19h-1.5A3.75 3.75 0 0 1 5 15.25v1.5Zm1.945 3.382l10.401-18.5l1.308.736l-10.402 18.5l-1.307-.736Z"/></g>`),
		g.Group(children),
	)
}